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

func DefaultFormatter(scalar float64, units string) string {
	return strings.TrimSpace(fmt.Sprintf("%v %v", strconv.FormatFloat(scalar, 'f', -1, 64), units))
}

func (q *Qty) Format(fn func(scalar float64, units string) string) string {
	return fn(q.scalar, q.Units())
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

	// using 2 slices since map doesn't have a defined iteration order
	var k []string
	var v []int

	for _, unit := range units {
		if i := slices.Index(k, unit); i >= 0 {
			v[i]++
		} else {
			k = append(k, unit)
			v = append(v, 1)
		}
	}

	result := []string{}
	for i := 0; i < len(k); i++ {
		if v[i] > 1 {
			result = append(result, fmt.Sprintf("%v^%v", k[i], v[i]))
		} else {
			result = append(result, k[i])
		}
	}
	return result
}
