package goqty

import (
	"math"
	"slices"
)

var signatureTypes = []string{"length", "time", "temperature", "mass", "current", "substance", "luminosity", "currency", "information", "angle"}

func (q *Qty) unitSignature() int {
	if q.signature != 0 && !q.IsUnitless() {
		return q.signature
	}

	vector := q.unitSignatureVector()
	for i, _ := range signatureTypes {
		vector[i] *= int(math.Pow(20, float64(i)))
	}

	return reduce(vector, func(prev, curr int) int {
		return prev + curr
	}, 0)
}

func (q *Qty) unitSignatureVector() []int {
	if !q.IsBase() {
		return q.ToBase().unitSignatureVector()
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
	return result
}
