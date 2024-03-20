package goqty

import (
	"regexp"
	"testing"
)

func Test(t *testing.T) {
	sign := "[+-]"
	integer := "\\d+"
	signedInteger := sign + "?" + integer
	fraction := "\\." + integer
	float := "(?:" + integer + "(?:" + fraction + ")?" + ")" +
		"|" +
		"(?:" + fraction + ")"
	exponent := "[Ee]" + signedInteger
	sciNumber := "(?:" + float + ")(?:" + exponent + ")?"

	signedNumber := sign + "?\\s*" + sciNumber
	qtyString := "(" + signedNumber + ")?" + "\\s*([^/]*)(?:/(.+))?"

	qtyStringRegex := regexp.MustCompile("^" + qtyString + "$")

	expr := "15.5 m/s^2"
	qtyMatches := qtyStringRegex.FindStringSubmatch(expr)

	t.Logf("%#v", qtyMatches)

	powerOp := "\\^|\\*{2}"
	safePower := "[01234]"
	topRegex := regexp.MustCompile("([^ \\*\\d]+?)(?:" + powerOp + ")?(-?" + safePower + ")")
	//bottomRegex := regexp.MustCompile("([^ \\*\\d]+?)(?:" + powerOp + ")?(" + safePower + "(?![a-zA-Z]))")

	top := "kilogram*meter*second^-2"
	matches := topRegex.FindStringSubmatch(top)

	t.Logf("%#v", matches)

	t.Fail()

}
