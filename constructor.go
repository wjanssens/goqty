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

	result.numerator = unityArray
	result.denominator = unityArray

	if units != "" {
		if q, err := ParseQty(units); err != nil {
			return result, err
		} else {
			result.numerator = q.numerator
			result.denominator = q.denominator
		}
	}

	// math with temperatures is very limited
	if strings.Contains(strings.Join(result.denominator, "*"), "temp") {
		return result, fmt.Errorf("cannot divide with temperatures")
	}
	if strings.Contains(strings.Join(result.numerator, "*"), "temp") {
		if len(result.numerator) > 1 {
			return result, fmt.Errorf("cannot divide with temperatures")
		}
		if slices.Compare(result.denominator, unityArray) != 0 {
			return result, fmt.Errorf("cannot divide with temperatures")
		}
	}

	result.updateBaseScalar()

	if result.IsTemperature() && result.baseScalar < 0 {
		return result, fmt.Errorf("temperatures must not be less than absolute zero")
	}
	return result, nil
}

func (q *Qty) Scalar() float64 {
	return q.scalar
}
func (q *Qty) Numerator() []string {
	return q.numerator
}
func (q *Qty) Denominator() []string {
	return q.denominator
}

func (q *Qty) updateBaseScalar() error {
	// if q.baseScalar != 0 {
	//   return q.baseScalar
	// }
	if q.IsBase() {
		q.baseScalar = q.scalar
		if signature, err := q.unitSignature(); err != nil {
			return err
		} else {
			q.signature = signature
		}
	} else {
		if base, err := q.ToBase(); err != nil {
			return err
		} else {
			q.baseScalar = base.scalar
			q.signature = base.signature
		}
	}
	return nil
}
