package goqty

import (
	"regexp"
	"slices"
)

func (q *Qty) IsUnitless() bool {
	return slices.Compare(q.numerator, unityArray) == 0 &&
		slices.Compare(q.denominator, unityArray) == 0
}

func (q *Qty) IsCompatible(other Qty) bool {
	return q.signature == other.signature
}

func (q *Qty) IsInverse(other Qty) bool {
	if i, err := q.Inverse(); err != nil {
		return false
	} else {
		return i.IsCompatible(other)
	}
}

func (q *Qty) IsBase() bool {
	if q.isBase != 0 {
		return q.isBase == 1
	}

	if q.IsDegrees() && regexp.MustCompile("<(kelvin|temp-K)>").MatchString(q.numerator[0]) {
		q.isBase = 1
		return true
	}

	units := slices.Concat(q.numerator, q.denominator)
	for _, u := range units {
		if u != unity && !slices.Contains(baseUnits, u) {
			q.isBase = -1
		}
	}
	if q.isBase == -1 {
		return false
	}
	q.isBase = 1
	return true
}
