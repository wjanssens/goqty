package goqty

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync"
)

var stringifiedUnitsCache sync.Map

func (q *Qty) Units() string {
	if q.units != "" {
		return q.units
	}

	numIsUnity := slices.Compare(q.numerator, unityArray) == 0
	denIsUnity := slices.Compare(q.denominator, unityArray) == 0
	if numIsUnity && denIsUnity {
		q.units = ""
		return q.units
	}

	var numUnits = StringifyUnits(q.numerator)
	var denUnits = StringifyUnits(q.denominator)
	if denIsUnity {
		q.units = numUnits
	} else {
		q.units = numUnits + "/" + denUnits
	}
	return q.units
}

func (q *Qty) String() string {
	return DefaultFormatter(q.scalar, q.Units())
}

type Formatter func(scalar float64, units string) string

func DefaultFormatter(scalar float64, units string) string {
	return strings.TrimSpace(fmt.Sprintf("%v %v", strconv.FormatFloat(scalar, 'f', -1, 64), units))
}

type FormattingOptions struct {
	TargetUnits *string
	Formatter   *Formatter
	Precision   int
}

func (q *Qty) Format(opts *FormattingOptions) (string, error) {
	if opts == nil {
		opts = &FormattingOptions{nil, nil, -1}
	}
	var target *Qty
	var err error
	if opts.TargetUnits != nil {
		if target, err = q.To(*opts.TargetUnits); target != nil {
			return "", err
		}
	} else {
		target = q
	}
	if opts.Formatter != nil {
		fn := *opts.Formatter
		return fn(target.scalar, target.Units()), nil
	} else {
		return strings.TrimSpace(fmt.Sprintf("%v %v", strconv.FormatFloat(target.scalar, 'f', -1, 64), target.Units())), nil
	}
}

func StringifyUnits(units []string) string {
	key := strings.Join(units, "|")
	if cached, found := stringifiedUnitsCache.Load(key); found {
		return cached.(string)
	}
	if isUnity := slices.Equal(units, unityArray); isUnity {
		stringifiedUnitsCache.Store(key, "1")
		return "1"
	} else {
		result := strings.Join(simplify(getOutputNames(units)), "*")
		stringifiedUnitsCache.Store(key, result)
		return result
	}
}

func getOutputNames(units []string) []string {
	result := []string{}
	for i := 0; i < len(units); i++ {
		token := units[i]
		if _, ok := prefixes[token]; ok {
			tokenNext := units[i+1]
			result = append(result, outputs[token]+outputs[tokenNext])
			i++
		} else {
			result = append(result, outputs[token])
		}
	}
	return result
}

func simplify(units []string) []string {
	// this turns ['s','m','s'] into ['s2','m']

	unitCounts := make(map[string]int)
	for _, unit := range units {
		if ct, ok := unitCounts[unit]; !ok {
			unitCounts[unit] = 1
		} else {
			unitCounts[unit] = ct + 1
		}
	}
	result := []string{}
	for k, v := range unitCounts {
		if v > 1 {
			result = append(result, fmt.Sprintf("%v%v", k, v))
		} else {
			result = append(result, k)
		}
	}
	return result
}
