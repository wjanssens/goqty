package goqty

import (
	"fmt"
	"slices"
	"strings"
	"sync"
)

var stringifiedUnitsCache sync.Map

func (q *Qty) Units() string {
	if q.units == "" {
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

func StringifyUnits(units []string) string {
	if cached, found := parsedUnitsCache.Load(units); found {
		return cached.(string)
	}
	if isUnity := slices.Compare(units, unityArray) == 0; isUnity {
		parsedUnitsCache.Store(units, "1")
		return "1"
	} else {
		result := strings.Join(simplify(getOutputNames(units)), "*")
		parsedUnitsCache.Store(units, result)
		return result
	}
}

func getOutputNames(units []string) []string {
	result := []string{}
	for i := 0; i < len(units); i++ {
		token := units[i]
		tokenNext := units[i+1]
		if _, ok := prefixes[token]; ok {
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
			unitCounts[unit] = 0
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
