package goqty

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		a        string
		b        string
		expected string
	}{
		"2.5m + 3m":  {"2.5m", "3m", "5.5 m"},
		"2.5m + 3cm": {"2.5m", "3cm", "2.53 m"},
		"3cm + 2.5m": {"3cm", "2.5m", "253 cm"},
		"3cm + 5cm":  {"3cm", "5cm", "8 cm"},
		"3m + 2s":    {"3m", "2s", "incompatible units m, s"},
		"10S + 0.1Ω": {"10S", "0.1Ω", "incompatible units S, Ω"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := ParseQty(test.a)
			if err != nil {
				t.Errorf("failed to parse %v, got %v", test.a, err)
				return
			}
			b, err := ParseQty(test.b)
			if err != nil {
				t.Errorf("failed to parse %v, got %v", test.b, err)
				return
			}
			if actual, err := a.Add(b); err != nil {
				if err.Error() != test.expected {
					t.Errorf("expected %v, got %v", test.expected, err)
				}
			} else {
				str := actual.String()
				if str != test.expected {
					t.Errorf("expected %v, got %v", test.expected, actual)
				}
			}
			if actual, err := a.Add(test.b); err != nil {
				if err.Error() != test.expected {
					t.Errorf("expected %v, got %v", test.expected, err)
				}
			} else {
				str := actual.String()
				if str != test.expected {
					t.Errorf("expected %v, got %v", test.expected, actual)
				}
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := map[string]struct {
		a        string
		b        string
		expected string
	}{
		"2.5m - 3m":  {"2.5m", "3m", "-0.5 m"},
		"3m - 2.5m":  {"3m", "2.5m", "0.5 m"},
		"2.5m - 3cm": {"2.5m", "3cm", "2.47 m"},
		"3cm - 2.5m": {"3cm", "2.5m", "-247 cm"},
		"3cm - 5cm":  {"3cm", "5cm", "-2 cm"},
		"3m - 2s":    {"3m", "2s", "incompatible units m, s"},
		"10S - 0.1Ω": {"10S", "0.1Ω", "incompatible units S, Ω"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := ParseQty(test.a)
			if err != nil {
				t.Errorf("failed to parse %v, got %v", test.a, err)
				return
			}
			b, err := ParseQty(test.b)
			if err != nil {
				t.Errorf("failed to parse %v got %v", test.b, err)
				return
			}
			if actual, err := a.Sub(b); err != nil {
				if err.Error() != test.expected {
					t.Errorf("expected %v, got %v", test.expected, err)
				}
			} else {
				str := actual.String()
				if str != test.expected {
					t.Errorf("expected %v, got %v", test.expected, actual)
				}
			}
			if actual, err := a.Sub(test.b); err != nil {
				if err.Error() != test.expected {
					t.Errorf("expected %v, got %v", test.expected, err)
				}
			} else {
				str := actual.String()
				if str != test.expected {
					t.Errorf("expected %v, got %v", test.expected, actual)
				}
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := map[string]struct {
		a        string
		b        string
		expected string
	}{
		"2.5m * 3m":  {"2.5m", "3m", "7.5 m2"},
		"3m * 2.5m":  {"3m", "2.5m", "7.5 m2"},
		"2.5m * 3cm": {"2.5m", "3cm", "0.075 m2"},
		"3cm * 2.5m": {"3cm", "2.5m", "750 cm2"},
		"2.5m * 3.5": {"2.5m", "3.5", "8.75 m"},
		"2.5m * 0.0": {"2.5m", "0.0", "0 m"},
		"10S * 0m":   {"2.5m", "0m", "0 m2"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := ParseQty(test.a)
			if err != nil {
				t.Errorf("failed to parse %v, got %v", test.a, err)
				return
			}
			b, err := ParseQty(test.b)
			if err != nil {
				t.Errorf("failed to parse %v got %v", test.b, err)
				return
			}
			if actual, err := a.Mul(b); err != nil {
				if err.Error() != test.expected {
					t.Errorf("expected %v, got %v", test.expected, err)
				}
			} else {
				str := actual.String()
				if str != test.expected {
					t.Errorf("expected %v, got %v", test.expected, actual)
				}
			}
			if actual, err := a.Mul(test.b); err != nil {
				if err.Error() != test.expected {
					t.Errorf("expected %v, got %v", test.expected, err)
				}
			} else {
				str := actual.String()
				if str != test.expected {
					t.Errorf("expected %v, got %v", test.expected, actual)
				}
			}
		})
	}
}
