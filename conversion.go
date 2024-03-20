package goqty

import (
	"fmt"
	"sync"
)

var conversionCache sync.Map

func (q *Qty) To(units string) (Qty, error) {
	if cached, found := conversionCache.Load(units); found {
		return cached.(Qty), nil
	}

	// Instantiating target to normalize units
	target, err := NewQty(1, units)
	if err != nil {
		return target, err
	} else if target.Units() == q.Units() {
		return *q, nil
	}

	if !q.IsCompatible(target) {
		if q.IsInverse(target) {
			i, _ := q.Inverse()
			if target, err = i.To(units); err != nil {
				return target, err
			}
		} else {
			return target, fmt.Errorf("Incompatible Units")
		}
	} else {
		if target.IsTemperature() {
			if target, err = ToTemp(*q, target); err != nil {
				return target, err
			}
		} else if target.IsDegrees() {
			if target, err = ToDegrees(*q, target); err != nil {
				return target, err
			}
		} else {
			if scalar, err := divSafe(q.baseScalar, target.baseScalar); err != nil {
				return target, nil
			} else {
				target = Qty{scalar: scalar, numerator: target.numerator, denominator: target.denominator}
			}
		}
	}

	conversionCache.Store(units, target)
	return target, nil
}
