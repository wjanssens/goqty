package goqty

import (
	"fmt"
	"math"
)

func filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

/**
 * Safely multiplies numbers while avoiding floating errors
 * like 0.1 * 0.1 => 0.010000000000000002
 *
 * @param {...number} numbers - numbers to multiply
 *
 * @returns {number} result
 */
func mulSafe(args ...float64) float64 {
	result := float64(1)
	decimals := float64(0)
	for _, a := range args {
		decimals += getFractional(a)
		result *= a
	}
	if decimals == 0 {
		return result
	} else {
		return round(result, decimals)
	}
}

/**
 * Safely divides two numbers while avoiding floating errors
 * like 0.3 / 0.05 => 5.999999999999999
 *
 * @returns {number} result
 * @param {number} num Numerator
 * @param {number} den Denominator
 */
func divSafe(num, den float64) (float64, error) {
	if den == 0 {
		return 0, fmt.Errorf("Divide by zero")
	}

	factor := math.Pow(10, getFractional(den))
	invDen := factor / (factor * den)

	return mulSafe(num, invDen), nil
}

// Rounds value at the specified number of decimals
func round(f, decimals float64) float64 {
	return math.Round(f*math.Pow(10, decimals)) / math.Pow(10, decimals)
}

func getFractional(f float64) float64 {
	if !isFinite(f) {
		return 0
	}
	count := 0
	for math.Mod(f, 1) != 0 {
		f *= 10
		count++
	}
	return float64(count)
}

func isFinite(f float64) bool {
	return !(math.IsInf(f, 0) || math.IsNaN(f))
}
