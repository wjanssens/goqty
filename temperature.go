package goqty

import (
	"fmt"
	"regexp"
	"slices"
)

var tempRegex = regexp.MustCompile("<temp-[CFRK]>")
var tempUnitRegex = regexp.MustCompile("<(kelvin|celsius|rankine|fahrenheit)>")

func (q *Qty) IsDegrees() bool {
	// signature may not have been calculated yet
	return (q.signature == 0 || q.signature == 400) &&
		len(q.numerator) == 1 &&
		slices.Equal(q.denominator, unityArray) &&
		(tempRegex.MatchString(q.numerator[0]) || tempUnitRegex.MatchString(q.numerator[0]))
}

func (q *Qty) IsTemperature() bool {
	return q.IsDegrees() && tempRegex.MatchString(q.numerator[0])
}

func subtractTemperatures(lhs, rhs *Qty) (*Qty, error) {
	lhsUnits := lhs.Units()
	rhsConverted, err := rhs.To(lhsUnits)
	if err != nil {
		return nil, err
	}
	dstDegreeUnits, err := getDegreeUnits(lhsUnits)
	if err != nil {
		return nil, err
	}
	dstDegrees, err := ParseQty(dstDegreeUnits)
	if err != nil {
		return nil, nil
	}
	return newQty(lhs.scalar-rhsConverted.scalar, dstDegrees.numerator, dstDegrees.denominator)
}

func subtractTempDegrees(temp, deg *Qty) (*Qty, error) {
	if units, err := getDegreeUnits(temp.Units()); err != nil {
		return nil, err
	} else {
		if tempDegrees, err := deg.To(units); err != nil {
			return nil, err
		} else {
			return newQty(temp.scalar-tempDegrees.scalar, temp.numerator, temp.denominator)
		}
	}
}

func addTempDegrees(temp, deg *Qty) (*Qty, error) {
	if units, err := getDegreeUnits(temp.Units()); err != nil {
		return nil, err
	} else {
		if tempDegrees, err := deg.To(units); err != nil {
			return nil, err
		} else {
			return newQty(temp.scalar+tempDegrees.scalar, temp.numerator, temp.denominator)
		}
	}
}

func getDegreeUnits(units string) (string, error) {
	switch units {
	case "tempK":
		return "degK", nil
	case "tempC":
		return "degC", nil
	case "tempF":
		return "degF", nil
	case "tempR":
		return "degR", nil
	default:
		return "", fmt.Errorf("unknown type for temp conversion from: %v", units)
	}
}

func ToDegrees(src, dst *Qty) (*Qty, error) {
	srcDegK, err := src.ToDegK()
	if err != nil {
		return nil, err
	}
	dstUnits := dst.Units()

	switch dstUnits {
	case "degK":
		dst.scalar = srcDegK.scalar
	case "degC":
		dst.scalar = srcDegK.scalar
	case "degF":
		dst.scalar = srcDegK.scalar * 9 / 5
	case "degR":
		dst.scalar = srcDegK.scalar * 9 / 5
	default:
		return nil, fmt.Errorf("unknown type for degree conversion to: %v", dstUnits)
	}
	return newQty(dst.scalar, dst.numerator, dst.denominator)
}

func (q *Qty) ToDegK() (*Qty, error) {
	var scalar float64
	units := q.Units()
	re := regexp.MustCompile("(deg)[CFRK]")
	if re.MatchString(units) {
		scalar = q.baseScalar
	} else {
		switch units {
		case "tempK":
			scalar = q.scalar
		case "tempC":
			scalar = q.scalar
		case "tempF":
			scalar = q.scalar * 5 / 9
		case "tempR":
			scalar = q.scalar * 5 / 9
		default:
			return nil, fmt.Errorf("unknown type for temp conversion from: %v", units)
		}
	}
	return newQty(scalar, []string{"<kelvin>"}, unityArray)
}

func ToTemp(src, dst *Qty) (*Qty, error) {
	dstUnits := dst.Units()
	var scalar float64
	switch dstUnits {
	case "tempK":
		scalar = src.baseScalar
	case "tempC":
		scalar = src.baseScalar - 273.15
	case "tempF":
		scalar = (src.baseScalar * 9.0 / 5.0) - 459.67
	case "tempR":
		scalar = src.baseScalar * 9.0 / 5.0
	default:
		return nil, fmt.Errorf("unknown type for temp conversion to: %v", dstUnits)
	}
	return newQty(scalar, dst.numerator, dst.denominator)
}

func (q *Qty) ToTempK() (*Qty, error) {
	units := q.Units()
	var scalar float64
	re := regexp.MustCompile("(deg)[CFRK]")
	if re.MatchString(units) {
		scalar = q.baseScalar
	} else {
		switch units {
		case "tempK":
			scalar = q.scalar
		case "tempC":
			scalar = q.scalar + 273.15
		case "tempF":
			scalar = (q.scalar + 459.67) * 5 / 9
		case "tempR":
			scalar = q.scalar * 5 / 9
		default:
			return nil, fmt.Errorf("unknown type for temp conversion from: %v", units)
		}
	}
	return newQty(scalar, []string{"<temp-K>"}, unityArray)
}
