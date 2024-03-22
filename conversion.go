package goqty

import (
	"fmt"
	"math"
	"sync"
)

var conversionCache sync.Map
var baseUnitCache sync.Map

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

// convert to base SI units
// results of the conversion are cached so subsequent calls to this will be fast
func (q *Qty) ToBase() (Qty, error) {
	if q.IsBase() {
		return *q, nil
	}
	if q.IsTemperature() {
		return q.ToTempK()
	}

	units := q.Units()
	if cached, found := baseUnitCache.Load(units); found {
		return cached.Mul(q.scalar)
	} else {
		cached = toBaseUnits(q.numerator, q.denominator)
		baseUnitCache.Store(units, cached)
		return cached.Mul(q.scalar)
	}
}

// Converts the unit back to a float if it is unitless.  Otherwise raises an exception
func (q *Qty) ToFloat() (float64, error) {
	if q.IsUnitless() {
		return q.scalar, nil
	} else {
		return q.scalar, fmt.Errorf("Can't convert to float unless unitless.  use Scalar()")
	}
}

/**
 * Returns the nearest multiple of quantity passed as
 * precision
 *
 * @param {(Qty|string|number)} precQuantity - Quantity, string formated
 *   quantity or number as expected precision
 *
 * @returns {Qty} Nearest multiple of precQuantity
 *
 * @example
 * Qty('5.5 ft').toPrec('2 ft'); // returns 6 ft
 * Qty('0.8 cu').toPrec('0.25 cu'); // returns 0.75 cu
 * Qty('6.3782 m').toPrec('cm'); // returns 6.38 m
 * Qty('1.146 MPa').toPrec('0.1 bar'); // returns 1.15 MPa
 *
 */
func (q *Qty) ToPrec(precision Qty) (Qty, error) {
	var err error
	if !q.IsUnitless() {
		if precision, err = precision.To(q.Units()); err != nil {
			return *q, err
		}
	} else if !precision.IsUnitless() {
		return *q, fmt.Errorf("Incompatible Units %v, %v", q.Units(), precision.Units())
	}

	if precision.scalar == 0 {
		return *q, fmt.Errorf("Divide by zero")
	}

	resultScalar := mulSafe(math.Round(q.scalar/precision.scalar), precision.scalar)

	return NewQty(resultScalar, q.Units())
}

/**
 * Configures and returns a fast function to convert
 * Number values from units to others.
 * Useful to efficiently convert large array of values
 * with same units into others with iterative methods.
 * Does not take care of rounding issues.
 *
 * @param {string} srcUnits Units of values to convert
 * @param {string} dstUnits Units to convert to
 *
 * @returns {Function} Converting function accepting Number value
 *   and returning converted value
 *
 * @throws "Incompatible units" if units are incompatible
 *
 * @example
 * // Converting large array of numbers with the same units
 * // into other units
 * var converter = Qty.swiftConverter("m/h", "ft/s");
 * var convertedSerie = largeSerie.map(converter);
 *
 */
func swiftConverter(srcUnits, dstUnits string) {
	srcQty := NewQty(srcUnits)
	dstQty := NewQty(dstUnits)

	if srcQty.Eq(dstQty) {
		return identity
	}

	//   var convert;
	//   if (!srcQty.isTemperature()) {
	//     convert = function(value) {
	//       return value * srcQty.baseScalar / dstQty.baseScalar;
	//     };
	//   }
	//   else {
	//     convert = function(value) {
	//       // TODO Not optimized
	//       return srcQty.mul(value).to(dstQty).scalar;
	//     };
	//   }

	//	return function converter(value) {
	//	  var i, length, result;
	//	  if (!Array.isArray(value)) {
	//	    return convert(value);
	//	  }
	//	  else {
	//	    length = value.length;
	//	    result = [];
	//	    for (i = 0; i < length; i++) {
	//	      result.push(convert(value[i]));
	//	    }
	//	    return result;
	//	  }
	//	};
}

func toBaseUnits(numerator, denominator []string) Qty {
	num := []string{}
	den := []string{}
	q := float64(1)

	for _, n := range numerator {
		if prefix, ok := prefixes[n]; ok {
			// workaround to fix
			// 0.1 * 0.1 => 0.010000000000000002
			q = mulSafe(q, prefix.scalar)
		} else if unit, ok := units[n]; ok {
			q *= unit.scalar
			num = append(num, unit.numerator...)
			den = append(den, unit.denominator...)
		}
	}

	for _, d := range denominator {
		if prefix, ok := prefixes[d]; ok {
			q /= prefix.scalar
		} else if unit, ok := units[d]; ok {
			q /= unit.scalar
			num = append(num, unit.numerator...)
			den = append(den, unit.denominator...)
		}

	}

	// num = reduce(num, func(a, b) {
	// 	return slices.Concat(a, b)
	// }, []string{})
	// den = reduce(den, func(a, b) {
	// 	return slices.Concat(a,b )
	// }, []string{})

	return Qty{scalar: q, numerator: num, denominator: den}
}
