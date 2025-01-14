package goqty

import (
	"slices"
	"testing"
)

func TestInit(t *testing.T) {
	qty, err := Parse("m")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(qty.numerator) != 1 || qty.numerator[0] != "<meter>" {
		t.Errorf("got %q, wanted %q", qty.numerator[0], "<meter>")
	}
	if qty.scalar != 1 {
		t.Errorf("got %v, wanted %v", qty.scalar, 1)
	}
}

func TestUnitless(t *testing.T) {
	qty, err := Parse("1")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if qty.scalar != 1 {
		t.Errorf("got %v, wanted %v", qty.scalar, 1.0)
	}
	if !slices.Equal(qty.numerator, []string{"<1>"}) {
		t.Errorf("got %v, wanted %v", qty.numerator, "[<1>]")
	}
	if !slices.Equal(qty.denominator, []string{"<1>"}) {
		t.Errorf("got %v, wanted %v", qty.denominator, "[<1>]")
	}
	qty, err = Parse("1.5")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if qty.scalar != 1.5 {
		t.Errorf("got %v, wanted %v", qty.scalar, 1.5)
	}
	if !slices.Equal(qty.numerator, []string{"<1>"}) {
		t.Errorf("got %v, wanted %v", qty.numerator, "[<1>]")
	}
	if !slices.Equal(qty.denominator, []string{"<1>"}) {
		t.Errorf("got %v, wanted %v", qty.denominator, "[<1>]")
	}
}
