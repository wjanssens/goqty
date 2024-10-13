package goqty

import (
	"fmt"
)

func (q *Qty) Add(other *Qty) (*Qty, error) {
	if !q.IsCompatible(other) {
		return nil, fmt.Errorf("incompatible Units %v, %v", q.Units(), other.Units())
	}
	if q.IsTemperature() && other.IsTemperature() {
		return nil, fmt.Errorf("cannot add two temperatures")
	} else if q.IsTemperature() {
		return addTempDegrees(q, other)
	} else if other.IsTemperature() {
		return addTempDegrees(other, q)
	}
	if to, err := other.To(q.Units()); err != nil {
		return nil, err
	} else {
		return newQty(q.scalar+to.scalar, q.numerator, q.denominator)
	}
}

func (q *Qty) Sub(other *Qty) (*Qty, error) {
	if !q.IsCompatible(other) {
		return nil, fmt.Errorf("incompatible Units %v, %v", q.Units(), other.Units())
	}

	if q.IsTemperature() && other.IsTemperature() {
		return subtractTemperatures(q, other)
	} else if q.IsTemperature() {
		return subtractTempDegrees(q, other)
	} else if other.IsTemperature() {
		return nil, fmt.Errorf("cannot subtract a temperature from a differential degree unit")
	}

	if to, err := other.To(q.units); err != nil {
		return nil, err
	} else {
		return newQty(q.scalar-to.scalar, q.numerator, q.denominator)
	}
}

func (q *Qty) Mul(input interface{}) (*Qty, error) {
	var other *Qty
	var err error
	switch t := input.(type) {
	case float64:
		return newQty(mulSafe(input.(float64), q.scalar), q.numerator, q.denominator)
	case float32:
		return newQty(mulSafe(float64(input.(float32)), q.scalar), q.numerator, q.denominator)
	case int:
		return newQty(mulSafe(float64(input.(int)), q.scalar), q.numerator, q.denominator)
	case Qty:
		cast := input.(Qty)
		other = &cast
	case *Qty:
		other = input.(*Qty)
	case string:
		if other, err = ParseQty(input.(string)); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("cannot multiply type %T", t)
	}

	if (q.IsTemperature() || other.IsTemperature()) && !(q.IsUnitless() || other.IsUnitless()) {
		return nil, fmt.Errorf("cannot multiply by temperatures")
	}

	// Quantities should be multiplied with same units if compatible, with base units else
	op1 := q
	op2 := other

	// so as not to confuse results, multiplication and division between temperature degrees will maintain original unit info in num/den
	// multiplication and division between deg[CFRK] can never factor each other out, only themselves: "degK*degC/degC^2" == "degK/degC"
	if op1.IsCompatible(op2) && op1.signature != 400 {
		if op2, err = op2.To(op1.units); err != nil {
			return nil, err
		}
	}
	if num, den, scale, err := cleanTerms(op1.numerator, op1.denominator, op2.numerator, op2.denominator); err != nil {
		return nil, err
	} else {
		scalar := mulSafe(op1.scalar, op2.scalar, scale)
		return newQty(scalar, num, den)
	}
}

func (q *Qty) Div(input interface{}) (*Qty, error) {
	var other *Qty
	var err error
	switch t := input.(type) {
	case float64:
		scalar := input.(float64)
		if scalar == 0.0 {
			return nil, fmt.Errorf("divide by zero")
		} else {
			return newQty(q.scalar/scalar, q.numerator, q.denominator)
		}
	case float32:
		scalar := float64(input.(float32))
		if scalar == 0 {
			return nil, fmt.Errorf("divide by zero")
		} else {
			return newQty(q.scalar/scalar, q.numerator, q.denominator)
		}
	case int:
		scalar := float64(input.(int))
		if scalar == 0 {
			return nil, fmt.Errorf("divide by zero")
		} else {
			return newQty(q.scalar/scalar, q.numerator, q.denominator)
		}
	case Qty:
		cast := input.(Qty)
		other = &cast
	case *Qty:
		other = input.(*Qty)
	case string:
		if other, err = ParseQty(input.(string)); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("cannot multiply type %T", t)
	}

	if other.scalar == 0 {
		return nil, fmt.Errorf("divide by zero")
	}

	if other.IsTemperature() {
		return nil, fmt.Errorf("cannot divide with temperatures")
	} else if q.IsTemperature() && !other.IsUnitless() {
		return nil, fmt.Errorf("cannot divide with temperatures")
	}

	// Quantities should be multiplied with same units if compatible, with base units else
	op1 := q
	op2 := other

	// so as not to confuse results, multiplication and division between temperature degrees will maintain original unit info in num/den
	// multiplication and division between deg[CFRK] can never factor each other out, only themselves: "degK*degC/degC^2" == "degK/degC"
	if op1.IsCompatible(op2) && op1.signature != 400 {
		if op2, err = op2.To(op1.units); err != nil {
			return nil, err
		}
	}
	if num, den, scale, err := cleanTerms(op1.numerator, op1.denominator, op2.denominator, op2.numerator); err != nil {
		return nil, err
	} else {
		return newQty(mulSafe(op1.scalar, scale)/op2.scalar, num, den)
	}
}

// // Returns a Qty that is the inverse of this Qty,
func (q *Qty) Inverse() (*Qty, error) {
	if q.IsTemperature() {
		return nil, fmt.Errorf("cannot divide with temperatures")
	}
	if q.scalar == 0 {
		return nil, fmt.Errorf("divide by zero")
	}
	return newQty(1/q.scalar, q.denominator, q.numerator)
}

type combinedType struct {
	dir    int
	term   string
	prefix string
	v1     float64
	v2     float64
}

func cleanTerms(num1, den1, num2, den2 []string) (num []string, den []string, scale float64, err error) {
	notUnity := func(val string) bool {
		return val != unity
	}

	num1 = filter(num1, notUnity)
	num2 = filter(num2, notUnity)
	den1 = filter(den1, notUnity)
	den2 = filter(den2, notUnity)

	combined := make(map[string]combinedType)

	combineTerms := func(terms []string, direction int) {
		var k string
		var prefix string

		for i, term := range terms {
			if _, ok := prefixes[term]; ok {
				k = terms[i+1]
				prefix = term
			} else {
				k = term
				prefix = unity
			}
			if k != "" && k != unity {
				if c, ok := combined[k]; ok {
					c.dir += direction
					combinedPrefixValue := prefixes[c.prefix].scalar
					if v, err := divSafe(prefixes[prefix].scalar, combinedPrefixValue); err != nil {
						// TODO
					} else if c.dir == 1 {
						c.v1 *= v
					} else {
						c.v2 *= v
					}
				} else {
					combined[k] = combinedType{dir: direction, term: k, prefix: prefix, v1: 1.0, v2: 1.0}
				}
			}
		}
	}

	combineTerms(num1, 1)
	combineTerms(den1, -1)
	combineTerms(num2, 1)
	combineTerms(den2, -1)

	num = []string{}
	den = []string{}
	scale = float64(1)

	for _, v := range combined {
		if v.dir > 0 {
			for n := 0; n < v.dir; n++ {
				if v.prefix == unity {
					num = append(num, v.term)
				} else {
					num = append(num, v.prefix, v.term)
				}
			}
		} else if v.dir < 0 {
			for n := 0; n < -v.dir; n++ {
				if v.prefix == unity {
					den = append(den, v.term)
				} else {
					den = append(den, v.prefix, v.term)
				}
			}
		}
		if s, err := divSafe(v.v1, v.v2); err != nil {
			return nil, nil, 0, err
		} else {
			scale *= s
		}
	}

	if len(num) == 0 {
		num = unityArray
	}
	if len(den) == 0 {
		den = unityArray
	}

	return num, den, scale, nil
}
