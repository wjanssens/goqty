package goqty

import (
	"math"
	"testing"
)

func TestTemperatureBaseUnit(t *testing.T) {
	qty, err := NewQty(1, "tempK")
	if err != nil {
		t.Errorf("failed to create '1 tempK'")
	}
	base, err := qty.ToBase()
	if err != nil {
		t.Errorf("failed to convert to base")
	}
	if base.scalar != 1 {
		t.Errorf("expected scalar %v, got %v", 1, base.scalar)
	}
	if base.Units() != "tempK" {
		t.Errorf("expected units %v, got %v", "tempK", base.Units())
	}

	qty, err = NewQty(1, "tempR")
	if err != nil {
		t.Errorf("failed to create '1 tempR'")
	}
	base, err = qty.ToBase()
	if err != nil {
		t.Errorf("failed to convert to base")
	}
	if base.scalar != 5.0/9.0 {
		t.Errorf("expected scalar %v, got %v", 5.0/9.0, base.scalar)
	}
	if base.Units() != "tempK" {
		t.Errorf("expected units %v, got %v", "tempK", base.Units())
	}

	qty, err = NewQty(0, "tempC")
	if err != nil {
		t.Errorf("failed to create '0 tempC'")
	}
	base, err = qty.ToBase()
	if err != nil {
		t.Errorf("failed to convert to base")
	}
	if base.scalar != 273.15 {
		t.Errorf("expected scalar %v, got %v", 273.15, base.scalar)
	}
	if base.Units() != "tempK" {
		t.Errorf("expected units %v, got %v", "tempK", base.Units())
	}

	qty, err = NewQty(0, "tempF")
	if err != nil {
		t.Errorf("failed to create '0 tempF'")
	}
	base, err = qty.ToBase()
	if err != nil {
		t.Errorf("failed to convert to base")
	}
	if math.Abs(base.scalar-255.372) > 0.001 {
		t.Errorf("expected scalar %v, got %v", 255.372, base.scalar)
	}
	if base.Units() != "tempK" {
		t.Errorf("expected units %v, got %v", "tempK", base.Units())
	}
}

func TestDegreesBaseUnit(t *testing.T) {
	qty, err := NewQty(1, "degK")
	if err != nil {
		t.Errorf("failed to create '1 degK'")
	}
	base, err := qty.ToBase()
	if err != nil {
		t.Errorf("failed to convert to base")
	}
	if base.scalar != 1 {
		t.Errorf("expected scalar %v, got %v", 1, base.scalar)
	}
	if base.Units() != "degK" {
		t.Errorf("expected units %v, got %v", "degK", base.Units())
	}

	qty, err = NewQty(1, "degR")
	if err != nil {
		t.Errorf("failed to create '1 degR'")
	}
	base, err = qty.ToBase()
	if err != nil {
		t.Errorf("failed to convert to base")
	}
	if base.scalar != 5.0/9.0 {
		t.Errorf("expected scalar %v, got %v", 5.0/9.0, base.scalar)
	}
	if base.Units() != "degK" {
		t.Errorf("expected units %v, got %v", "degK", base.Units())
	}

	qty, err = NewQty(0, "degC")
	if err != nil {
		t.Errorf("failed to create '0 degC'")
	}
	base, err = qty.ToBase()
	if err != nil {
		t.Errorf("failed to convert to base")
	}
	if base.scalar != 273.15 {
		t.Errorf("expected scalar %v, got %v", 273.15, base.scalar)
	}
	if base.Units() != "degK" {
		t.Errorf("expected units %v, got %v", "degK", base.Units())
	}

	qty, err = NewQty(0, "degF")
	if err != nil {
		t.Errorf("failed to create '0 degF'")
	}
	base, err = qty.ToBase()
	if err != nil {
		t.Errorf("failed to convert to base")
	}
	if base.scalar != 5/9 {
		t.Errorf("expected scalar %v, got %v", 5/9, base.scalar)
	}
	if base.Units() != "degK" {
		t.Errorf("expected units %v, got %v", "degK", base.Units())
	}
}
