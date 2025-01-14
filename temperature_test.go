package goqty

import (
	"math"
	"testing"
)

func TestTemperatureBaseUnit(t *testing.T) {
	tests := map[string]struct {
		q      string
		scalar float64
		unit   string
	}{
		"1 tempK": {"1 tempK", 1, "tempK"},
		"1 tempR": {"1 tempR", 5.0 / 9.0, "tempK"},
		"0 tempC": {"0 tempC", 273.15, "tempK"},
		"0 tempF": {"0 tempF", 255.372, "tempK"},
		"1 degK":  {"1 degK", 1, "\u00b0K"},
		"1 degR":  {"1 degR", 5.0 / 9.0, "\u00b0K"},
		"1 degC":  {"1 degC", 1, "\u00b0K"},
		"1 degF":  {"1 degF", 5.0 / 9.0, "\u00b0K"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			qty, err := Parse(test.q)
			if err != nil {
				t.Errorf("failed to create %v", test.q)
			}
			base, err := qty.ToBase()
			if err != nil {
				t.Errorf("failed to convert to base, got %v", err)
			} else {
				if math.Abs(base.scalar-test.scalar) > 0.001 {
					t.Errorf("expected scalar %v, got %v", test.scalar, base.scalar)
				}
				u := base.Units()
				if base.Units() != test.unit {
					t.Errorf("expected units %v, got %v", test.unit, u)
				}
			}
		})
	}
}

func TestAbsoluteZero(t *testing.T) {
	if q, err := Parse("-1 tempK"); err == nil {
		t.Errorf("expected error, got %v", q)
	}
	if q, err := Parse("-273.16 tempC"); err == nil {
		t.Errorf("expected error, got %v", q)
	}
	if q, err := Parse("-459.68 tempF"); err == nil {
		t.Errorf("expected error, got %v", q)
	}
	if q, err := Parse("-1 tempR"); err == nil {
		t.Errorf("expected error, got %v", q)
	}

	if q, err := Parse("1 tempK"); err != nil {
		t.Errorf("failed to create '1 tempK': %v", err)
	} else {
		if q, err := q.Mul(-1); err == nil {
			t.Errorf("expected error, got %v", q)
		}
	}
	if q, err := Parse("0 tempK"); err != nil {
		t.Errorf("failed to create '0 tempK': %v", err)
	} else {
		if q, err := q.Sub("1 degK"); err == nil {
			t.Errorf("expected error, got %v", q)
		}
	}
	if q, err := Parse("-273.15 tempC"); err != nil {
		t.Errorf("failed to create '0 tempK': %v", err)
	} else {
		if q, err := q.Sub("1 degC"); err == nil {
			t.Errorf("expected error, got %v", q)
		}
	}
	if q, err := Parse("-459.67 tempF"); err != nil {
		t.Errorf("failed to create '0 tempK': %v", err)
	} else {
		if q, err := q.Sub("1 degF"); err == nil {
			t.Errorf("expected error, got %v", q)
		}
	}
	if q, err := Parse("0 tempR"); err != nil {
		t.Errorf("failed to create '0 tempK': %v", err)
	} else {
		if q, err := q.Sub("1 degR"); err == nil {
			t.Errorf("expected error, got %v", q)
		}
	}

}
