package goqty

import (
	"slices"
	"testing"
)

func TestParseUnitOnly(t *testing.T) {
	qty, err := ParseQty("m")
	if err != nil {
		t.Errorf("failed to parse 'm'")
	}
	if slices.Equal(qty.numerator, []string{"<meter>"}) {
		t.Errorf("expected numerator %v, got %v", "<meter>", qty.numerator)
	}
	if qty.scalar != 1 {
		t.Errorf("expected scalar %v, got %v", 1, qty.scalar)
	}
}
func TestParseUnitless(t *testing.T) {
	qty, err := ParseQty("1")
	if err != nil {
		t.Errorf("failed to parse '1'")
	}
	float, err := qty.ToFloat()
	if err != nil {
		t.Errorf("failed to convert to float")
	}
	if float != 1 {
		t.Errorf("expected float %v, got %v", 1, float)
	}
	expected := []string{"<1>"}
	if !slices.Equal(qty.numerator, expected) {
		t.Errorf("expected numerator %v, got %v", expected, qty.numerator)
	}
	if !slices.Equal(qty.denominator, expected) {
		t.Errorf("expected denominator %v, got %v", expected, qty.denominator)
	}

	qty, err = ParseQty("1.5")
	if err != nil {
		t.Errorf("failed to parse '1.5'")
	}
	float, err = qty.ToFloat()
	if err != nil {
		t.Errorf("failed to convert to float")
	}
	if float != 1.5 {
		t.Errorf("expected float %v, got %v", 1.5, float)
	}
	if !slices.Equal(qty.numerator, expected) {
		t.Errorf("expected numerator %v, got %v", expected, qty.numerator)
	}
	if !slices.Equal(qty.denominator, expected) {
		t.Errorf("expected denominator %v, got %v", expected, qty.denominator)
	}
}

func TestNonASCIICharacter(t *testing.T) {
	var qty *Qty
	var expected *Qty
	var err error

	expected, err = ParseQty("1 um")
	if err != nil {
		t.Errorf("failed to parse '1 um', got %v", err)
	}

	// greek letter
	qty, err = ParseQty("1 \u03BCm")
	if err != nil {
		t.Errorf("failed to parse '1 \u03BCm', got %v", err)
	} else {
		if !qty.Eq(expected) {
			t.Errorf("expected %v, got %v", expected, qty)
		}
	}
	// micro sign
	qty, err = ParseQty("1 \u00B5m")
	if err != nil {
		t.Errorf("failed to parse '1 \u03BCm', got %v", err)
	} else {
		if !qty.Eq(expected) {
			t.Errorf("expected %v, got %v", expected, qty)
		}
	}
	expected, err = ParseQty("1 ohm")
	if err != nil {
		t.Errorf("failed to parse '1 ohm', got %v", err)
	}

	// greek letter
	qty, err = ParseQty("1 \u03A9")
	if err != nil {
		t.Errorf("failed to parse '1 \u03BCm', got %v", err)
	} else {
		if !qty.Eq(expected) {
			t.Errorf("expected %v, got %v", expected, qty)
		}
	}
	// ohm sign
	qty, err = ParseQty("1 \u2126")
	if err != nil {
		t.Errorf("failed to parse '1 \u03BCm', got %v", err)
	} else {
		if !qty.Eq(expected) {
			t.Errorf("expected %v, got %v", expected, qty)
		}
	}
}
