package goqty

import (
	"slices"
	"testing"
)

func TestNewUnitless(t *testing.T) {
	qty, err := NewQty(1.5, "")
	if err != nil {
		t.Errorf("failed to create '1.5'")
	}
	float, err := qty.ToFloat()
	if err != nil {
		t.Errorf("failed to convert to float: %v", err)
	}
	if float != 1.5 {
		t.Errorf("expected float %v, got %v", 1.5, float)
	}
	expected := []string{"<1>"}
	if !slices.Equal(qty.numerator, expected) {
		t.Errorf("expected numerator %v, got %v", expected, qty.numerator)
	}
	if !slices.Equal(qty.denominator, expected) {
		t.Errorf("expected denominator %v, got %v", expected, qty.denominator)
	}
}

func TestNewWithUnit(t *testing.T) {
	qty, err := NewQty(1.5, "m")
	if err != nil {
		t.Errorf("failed to create '1.5 m'")
	}
	if qty.scalar != 1.5 {
		t.Errorf("expected scalar %v, got %v", 1.5, float)
	}
	expected := []string{"<meter>"}
	if !slices.Equal(qty.numerator, expected) {
		t.Errorf("expected numerator %v, got %v", expected, qty.numerator)
	}
	expected = []string{"<1>"}
	if !slices.Equal(qty.denominator, expected) {
		t.Errorf("expected denominator %v, got %v", expected, qty.denominator)
	}
}
