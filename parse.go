package goqty

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// func Parse(expr string) Qty {

// }
// func NewQty(scalar float64, unit string) Qty {

// }

const sign = "[+-]"
const integer = "\\d+"
const signedInteger = sign + "?" + integer
const fraction = "\\." + integer
const float = "(?:" + integer + "(?:" + fraction + ")?" + ")" +
	"|" +
	"(?:" + fraction + ")"
const exponent = "[Ee]" + signedInteger
const sciNumber = "(?:" + float + ")(?:" + exponent + ")?"
const signedNumber = sign + "?\\s*" + sciNumber
const qtyString = "(" + signedNumber + ")?" + "\\s*([^/]*)(?:/(.+))?"

var qtyStringRegex = regexp.MustCompile("^" + qtyString + "$")

const powerOp = "\\^|\\*{2}"
const safePower = "[01234]"

var topRegex = regexp.MustCompile("([^ \\*\\d]+?)(?:" + powerOp + ")?(-?" + safePower + ")")
var bottomRegex = regexp.MustCompile("([^ \\*\\d]+?)(?:" + powerOp + ")?(" + safePower + ")")

var prefix = re(prefixesByAlias)
var unit = re(unitsByAlias)
var boundary = "\\b|$" // TODO \b only supports ASCII
var unitMatch = "(" + prefix + ")??(" + unit + ")(?:" + boundary + ")"
var unitTestRegex = regexp.MustCompile("^\\s*(" + unitMatch + "[\\s\\*]*)+$")

var wsRegex = regexp.MustCompile(`\\s`)

var parsedUnitsCache sync.Map

/* parse a string into a unit object.
 * Typical formats like :
 * "5.6 kg*m/s^2"
 * "5.6 kg*m*s^-2"
 * "5.6 kilogram*meter*second^-2"
 * "2.2 kPa"
 * "37 degC"
 * "1"  -- creates a unitless constant with value 1
 * "GPa"  -- creates a unit with scalar 1 with units 'GPa'
 * 6'4"  -- recognized as 6 feet + 4 inches
 * 8 lbs 8 oz -- recognized as 8 lbs + 8 ounces
 */
func ParseQty(expr string) (*Qty, error) {
	result := Qty{
		scalar:      1,
		numerator:   unityArray,
		denominator: unityArray,
	}

	expr = strings.TrimSpace(expr)
	qtyMatches := qtyStringRegex.FindStringSubmatch(expr)
	if qtyMatches == nil {
		return nil, fmt.Errorf("%v: Quantity not recognized", expr)
	}
	scalar := qtyMatches[1]
	top := qtyMatches[2]
	bottom := qtyMatches[3]

	if scalar != "" {
		// Allow whitespaces between sign and scalar for loose parsing
		scalarMatch := wsRegex.ReplaceAllString(scalar, "")
		result.scalar, _ = strconv.ParseFloat(scalarMatch, 64)
	} else {
		result.scalar = 1
	}

	var err error
	var n int64
	var x, nx string
	for {
		matches := topRegex.FindStringSubmatch(top)
		if matches == nil {
			break
		}
		unit := matches[1]
		power := matches[2]

		if n, err = strconv.ParseInt(power, 10, 8); err != nil {
			return nil, fmt.Errorf("unit exponenent is not a number")
		}
		if unitTestRegex.FindString(unit) == "" {
			return nil, fmt.Errorf("unit is not recognized")
		}
		x = unit + " "
		nx = ""
		for i := 0; i < int(n); i++ {
			nx += x
		}
		if n >= 0 {
			top = strings.ReplaceAll(top, matches[0], nx)
		} else {
			if bottom == "" {
				bottom = nx
			} else {
				bottom = bottom + nx
			}
			top = strings.ReplaceAll(top, matches[0], "")
		}
	}

	for {
		matches := bottomRegex.FindStringSubmatch(bottom)
		if matches == nil {
			break
		}
		unit := matches[1]
		power := matches[2]

		if n, err = strconv.ParseInt(power, 10, 8); err != nil {
			return nil, fmt.Errorf("unit exponenent is not a number")
		}
		if unitTestRegex.FindString(unit) == "" {
			return nil, fmt.Errorf("unit is not recognized")
		}
		x = unit + " "
		nx = ""
		for j := 0; j < int(n); j++ {
			nx += x
		}
		if n >= 0 {
			bottom = strings.ReplaceAll(bottom, matches[0], nx)
		} else {
			if bottom == "" {
				bottom = nx
			} else {
				bottom = bottom + nx
			}
			bottom = strings.ReplaceAll(bottom, matches[0], "")
		}
	}

	if top != "" {
		if result.numerator, err = parseUnits(strings.TrimSpace(top)); err != nil {
			return nil, fmt.Errorf("unparsable numerator units")
		}
	}
	if bottom != "" {
		if result.denominator, err = parseUnits(strings.TrimSpace(bottom)); err != nil {
			return nil, fmt.Errorf("unparsable denominator units")
		}
	}

	return &result, nil
}

/* Parses and convers units string to normalized units array.
 * Result is cached to speed up future calls.
 */
func parseUnits(units string) ([]string, error) {
	if cached, found := parsedUnitsCache.Load(units); found {
		return cached.([]string), nil
	}

	matches := unitTestRegex.FindAllStringSubmatch(units, -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("Unit not recognized")
	}
	result := make([]string, 0)
	for _, match := range matches {
		prefix, hasPrefix := prefixesByAlias[match[2]]
		unit, hasUnit := unitsByAlias[match[3]]

		if hasPrefix && hasUnit {
			result = append(result, prefix, unit)
		} else if hasUnit {
			result = append(result, unit)
		}
	}
	parsedUnitsCache.Store(units, result)
	return result, nil
}

func re(unitsByAlias map[string]string) string {
	keys := make([]string, len(unitsByAlias))
	i := 0
	for k := range unitsByAlias {
		keys[i] = k
		i++
	}
	sort.SliceStable(keys, func(i int, j int) bool {
		return len(keys[i]) < len(keys[j])
	})
	return strings.Join(keys, "|")
}
