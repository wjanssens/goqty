package qty

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

func newQty(scalar float64, numerator []string, denominator []string) (*Qty, error) {
	result := Qty{
		scalar:      scalar,
		numerator:   numerator,
		denominator: denominator,
	}
	if len(numerator) == 0 {
		result.numerator = unityArray
	}
	if len(denominator) == 0 {
		result.denominator = unityArray
	}

	// math with temperatures is very limited
	if strings.Contains(strings.Join(result.denominator, "*"), "temp") {
		return nil, fmt.Errorf("cannot divide with temperatures")
	}
	if strings.Contains(strings.Join(result.numerator, "*"), "temp") {
		if len(result.numerator) > 1 {
			return nil, fmt.Errorf("cannot divide with temperatures")
		}
		if slices.Compare(result.denominator, unityArray) != 0 {
			return nil, fmt.Errorf("cannot divide with temperatures")
		}
	}

	if err := result.updateBaseScalar(); err != nil {
		return nil, err
	}

	if result.IsTemperature() && result.baseScalar < 0 {
		return nil, fmt.Errorf("temperatures must not be less than absolute zero")
	}
	return &result, nil

}
func New(scalar float64, units string) (*Qty, error) {
	if units != "" {
		if q, err := Parse(units); err != nil {
			return nil, err
		} else {
			return newQty(scalar, q.numerator, q.denominator)
		}
	} else {
		return newQty(scalar, unityArray, unityArray)
	}
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
