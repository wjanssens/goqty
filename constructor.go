package goqty

import (
	"fmt"
	"slices"
	"strings"
)

type Qty struct {
	scalar      float64
	baseScalar  float64
	numerator   []string
	denominator []string
	units       string
	signature   int
	isBase      int
}

func NewQty(scalar float64, units string) (Qty, error) {
	result := Qty{
		scalar: scalar,
	}

	if units != "" {
		if q, err := Parse(units); err != nil {
			return result, err
		} else {
			result.numerator = q.numerator
			result.denominator = q.denominator
		}
	}

	// math with temperatures is very limited
	if strings.Index(strings.Join(result.denominator, "*"), "temp") >= 0 {
		return result, fmt.Errorf("Cannot divide with temperatures")
	}
	if strings.Index(strings.Join(result.numerator, "*"), "temp") >= 0 {
		if len(result.numerator) > 1 {
			return result, fmt.Errorf("Cannot divide with temperatures")
		}
		if slices.Compare(result.denominator, unityArray) != 0 {
			return result, fmt.Errorf("Cannot divide with temperatures")
		}
	}

	result.updateBaseScalar()

	if result.IsTemperature() && result.baseScalar < 0 {
		return result, fmt.Errorf("Temperatures must not be less than absolute zero")
	}
	return result, nil
}

func (q *Qty) updateBaseScalar() {
	// if q.baseScalar != 0 {
	//   return q.baseScalar
	// }
	if q.IsBase() {
		q.baseScalar = q.scalar
		q.signature = unitSignature.call(this)
	} else {
		var base = q.ToBase()
		q.baseScalar = base.scalar
		q.signature = base.signature
	}
}
