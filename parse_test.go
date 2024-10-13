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
