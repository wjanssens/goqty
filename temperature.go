package goqty

import (
	"fmt"
	"regexp"
	"slices"
)

var tempRegex = regexp.MustCompile("<temp-[CFRK]>")
var tempUnitRegex = regexp.MustCompile("<(kelvin|celsius|rankine|fahrenheit)>")

func (q *Qty) IsDegrees() bool {
	return (q.signature == 0 || q.signature == 400) &&
		len(q.numerator) == 1 &&
		slices.Compare(q.denominator, unityArray) == 0 &&
		(tempRegex.MatchString(q.numerator[0]) || tempUnitRegex.MatchString(q.numerator[0]))
}

func (q *Qty) IsTemperature() bool {
	return q.IsDegrees() && tempRegex.MatchString(q.numerator[0])
}

func subtractTemperatures(lhs, rhs Qty) (Qty, error) {
	lhsUnits := lhs.Units()
	rhsConverted, err := rhs.To(lhsUnits)
	if err != nil {
		return lhs, err
	}
	dstDegreeUnits, err := getDegreeUnits(lhsUnits)
	if err != nil {
		return lhs, err
	}
	dstDegrees, err := Parse(dstDegreeUnits)
	if err != nil {
		return lhs, nil
	}
	return Qty{scalar: lhs.scalar - rhsConverted.scalar, numerator: dstDegrees.numerator, denominator: dstDegrees.denominator}, nil
}

// func subtractTempDegrees(temp,deg) {
// 	var tempDegrees = deg.to(getDegreeUnits(temp.units()));
// 	return Qty({"scalar": temp.scalar - tempDegrees.scalar, "numerator": temp.numerator, "denominator": temp.denominator});
//   }

// func addTempDegrees(temp,deg) {
// 	var tempDegrees = deg.to(getDegreeUnits(temp.units()));
// 	return Qty({"scalar": temp.scalar + tempDegrees.scalar, "numerator": temp.numerator, "denominator": temp.denominator});
//   }

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
		return "", fmt.Errorf("Unknown type for temp conversion from: " + units)
	}
}

func ToDegrees(src, dst Qty) (Qty, error) {
	result := Qty{
		numerator:   dst.numerator,
		denominator: dst.denominator,
	}

	srcDegK, err := src.ToDegK()
	if err != nil {
		return result, err
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
		return result, fmt.Errorf("Unknown type for degree conversion to: " + dstUnits)
	}
	return result, nil
}

func (q *Qty) ToDegK() (Qty, error) {
	var units = q.Units()

	result := Qty{
		numerator:   []string{"<kelvin>"},
		denominator: unityArray,
	}
	re := regexp.MustCompile("(deg)[CFRK]")
	if re.MatchString(units) {
		result.scalar = q.baseScalar
	} else {
		switch units {
		case "tempK":
			result.scalar = q.scalar
		case "tempC":
			result.scalar = q.scalar
		case "tempF":
			result.scalar = q.scalar * 5 / 9
		case "tempR":
			result.scalar = q.scalar * 5 / 9
		default:
			return result, fmt.Errorf("Unknown type for temp conversion from: " + units)
		}
	}
	return result, nil
}

func ToTemp(src, dst Qty) (Qty, error) {
	dstUnits := dst.Units()

	result := Qty{
		numerator:   dst.numerator,
		denominator: dst.denominator,
	}
	switch dstUnits {
	case "tempK":
		dst.scalar = src.baseScalar
	case "tempC":
		dst.scalar = src.baseScalar - 273.15
	case "tempF":
		dst.scalar = (src.baseScalar * 9 / 5) - 459.67
	case "tempR":
		dst.scalar = src.baseScalar * 9 / 5
	default:
		return result, fmt.Errorf("Unknown type for temp conversion to: " + dstUnits)
	}
	return result, nil
}

func (q *Qty) ToTempK() (Qty, error) {
	var units = q.Units()

	result := Qty{
		numerator:   []string{"<kelvin>"},
		denominator: unityArray,
	}
	re := regexp.MustCompile("(deg)[CFRK]")
	if re.MatchString(units) {
		result.scalar = q.baseScalar
	} else {
		switch units {
		case "tempK":
			result.scalar = q.scalar
		case "tempC":
			result.scalar = q.scalar + 273.15
		case "tempF":
			result.scalar = (q.scalar + 459.67) * 5 / 9
		case "tempR":
			result.scalar = q.scalar * 5 / 9
		default:
			return result, fmt.Errorf("Unknown type for temp conversion from: " + units)
		}
	}
	return result, nil
}
