package goqty

import (
	"math"
	"slices"
)

var signatureTypes = []string{"length", "time", "temperature", "mass", "current", "substance", "luminosity", "currency", "information", "angle"}

// calculates the unit signature id for use in comparing compatible units and simplification
// the signature is based on a simple classification of units and is based on the following publication

// Novak, G.S., Jr. "Conversion of units of measurement", IEEE Transactions on Software Engineering,
// 21(8), Aug 1995, pp.651-661
// doi://10.1109/32.403789
// http://ieeexplore.ieee.org/Xplore/login.jsp?url=/iel1/32/9079/00403789.pdf?isnumber=9079&prod=JNL&arnumber=403789&arSt=651&ared=661&arAuthor=Novak%2C+G.S.%2C+Jr.
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

		result := reduce(vector, func(prev, curr int) int {
			return prev + curr
		}, 0)
		return result, nil
	}
}

// calculates the unit signature vector used by unit_signature
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
