package goqty

import (
	"slices"
	"testing"
)

func TestNewUnitless(t *testing.T) {
	qty, err := NewQty(1.5, "")
	if err != nil {
		t.Errorf("failed to create '1.5', got %v", err)
	}
	float, err := qty.ToFloat()
	if err != nil {
		t.Errorf("failed to convert to float, got %v", err)
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
		t.Errorf("failed to create '1.5 m', got %v", err)
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

func TestNewSimple(t *testing.T) {
	if qty, err := ParseQty("1m"); err != nil {
		t.Errorf("failed to create '1m', got %v", err)
	} else {
		if qty.scalar != 1.0 {
			t.Errorf("expected scalar %v, got %v", 1.0, qty.scalar)
		}
		expected := []string{"<meter>"}
		if !slices.Equal(qty.numerator, expected) {
			t.Errorf("expected numerator %v, got %v", expected, qty.numerator)
		}
		expected = []string{"<1>"}
		if !slices.Equal(qty.denominator, expected) {
			t.Errorf("expected numerator %v, got %v", expected, qty.denominator)
		}
	}
}

func TestNewNegative(t *testing.T) {
	if qty, err := ParseQty("-1m"); err != nil {
		t.Errorf("failed to create '1m', got %v", err)
	} else {
		if qty.scalar != -1.0 {
			t.Errorf("expected scalar %v, got %v", -1.0, qty.scalar)
		}
		expected := []string{"<meter>"}
		if !slices.Equal(qty.numerator, expected) {
			t.Errorf("expected numerator %v, got %v", expected, qty.numerator)
		}
		expected = []string{"<1>"}
		if !slices.Equal(qty.denominator, expected) {
			t.Errorf("expected numerator %v, got %v", expected, qty.denominator)
		}
	}
}

func TestNewCompound(t *testing.T) {
	if qty, err := ParseQty("-1 N*m"); err != nil {
		t.Errorf("failed to create '1 N*m', got %v", err)
	} else {
		if qty.scalar != -1.0 {
			t.Errorf("expected scalar %v, got %v", -1.0, qty.scalar)
		}
		expected := []string{"<newton>", "<meter>"}
		if !slices.Equal(qty.numerator, expected) {
			t.Errorf("expected numerator %v, got %v", expected, qty.numerator)
		}
		expected = []string{"<1>"}
		if !slices.Equal(qty.denominator, expected) {
			t.Errorf("expected numerator %v, got %v", expected, qty.denominator)
		}
	}
}

func TestPressureUnits(t *testing.T) {
	if qty, err := ParseQty("1 inH2O"); err != nil {
		t.Errorf("failed to create '1 inH2O, got %v", err)
	} else {
		if qty.scalar != 1 {
			t.Errorf("expected scalar %v, got %v", 1.0, qty.scalar)
		}
		expected := []string{"<inh2o>"}
		if !slices.Equal(qty.numerator, expected) {
			t.Errorf("expected numerator %v, got %v", expected, qty.numerator)
		}
	}

	if qty, err := ParseQty("1 cmH2O"); err != nil {
		t.Errorf("failed to create '1 cmH2O, got %v", err)
	} else {
		if qty.scalar != 1 {
			t.Errorf("expected scalar %v, got %v", 1.0, qty.scalar)
		}
		expected := []string{"<cmh2o>"}
		if !slices.Equal(qty.numerator, expected) {
			t.Errorf("expected numerator %v, got %v", expected, qty.numerator)
		}
	}
}

func TestNewWithDenominator(t *testing.T) {
	if qty, err := ParseQty("1 m/s"); err != nil {
		t.Errorf("failed to create '1 m/s, got %v", err)
	} else {
		if qty.scalar != 1 {
			t.Errorf("expected scalar %v, got %v", 1.0, qty.scalar)
		}
		expected := []string{"<meter>"}
		if !slices.Equal(qty.numerator, expected) {
			t.Errorf("expected numerator %v, got %v", expected, qty.numerator)
		}
		expected = []string{"<second>"}
		if !slices.Equal(qty.denominator, expected) {
			t.Errorf("expected denominator %v, got %v", expected, qty.denominator)
		}
	}
}
