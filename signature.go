package goqty

import (
	"math"
	"slices"
)

var signatureTypes = []string{"length", "time", "temperature", "mass", "current", "substance", "luminosity", "currency", "information", "angle"}

func (q *Qty) unitSignature() (int, error) {
	if q.signature != 0 && !q.IsUnitless() {
		return q.signature, nil
	}

	if vector, err := q.unitSignatureVector(); err != nil {
		return q.signature, err
	} else {
		for i, _ := range signatureTypes {
			vector[i] *= int(math.Pow(20, float64(i)))
		}

		return reduce(vector, func(prev, curr int) int {
			return prev + curr
		}, 0), nil
	}
}

func (q *Qty) unitSignatureVector() ([]int, error) {
	if !q.IsBase() {
		if b, err := q.ToBase(); err != nil {
			return []int{}, err
		} else {
			return b.unitSignatureVector()
		}
	}

	result := make([]int, len(signatureTypes))
	for i, _ := range result {
		result[i] = 0
	}
	for _, v := range q.numerator {
		if r, ok := units[v]; ok {
			if n := slices.Index(signatureTypes, r.kind); n >= 0 {
				result[n] = result[n] + 1
			}
		}
	}
	for _, v := range q.denominator {
		if r, ok := units[v]; ok {
			if n := slices.Index(signatureTypes, r.kind); n >= 0 {
				result[n] = result[n] - 1
			}
		}
	}
	return result, nil
}
