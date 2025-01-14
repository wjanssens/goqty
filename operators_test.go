package goqty

import (
	"testing"
)

func TestInverse(t *testing.T) {
	tests := map[string]struct {
		expr     string
		expected string
	}{
		"tempF": {"tempF", "cannot divide with temperatures"},
		"10S":   {"10 S", "0.1 1/S"},
		"0 ohm": {"0 ohm", "divide by zero"},
		"8 in":  {"8 in", "0.125 1/in"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := Parse(test.expr)
			if err != nil {
				t.Errorf("failed to parse %v, got %v", test.expr, err)
				return
			}
			if actual, err := a.Inverse(); err != nil {
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
		"3m + 2s":    {"3m", "2s", "incompatible units: m and s"},
		"10S + 0.1Ω": {"10S", "0.1Ω", "incompatible units: S and Ω"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := Parse(test.a)
			if err != nil {
				t.Errorf("failed to parse %v, got %v", test.a, err)
				return
			}
			b, err := Parse(test.b)
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
		"3m - 2s":    {"3m", "2s", "incompatible units: m and s"},
		"10S - 0.1Ω": {"10S", "0.1Ω", "incompatible units: S and Ω"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := Parse(test.a)
			if err != nil {
				t.Errorf("failed to parse %v, got %v", test.a, err)
				return
			}
			b, err := Parse(test.b)
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
		"2.5m * 3m":  {"2.5m", "3m", "7.5 m^2"},
		"3m * 2.5m":  {"3m", "2.5m", "7.5 m^2"},
		"2.5m * 3cm": {"2.5m", "3cm", "0.075 m^2"},
		"3cm * 2.5m": {"3cm", "2.5m", "750 cm^2"},
		// unitless
		"2.5m * 3.5": {"2.5m", "3.5", "8.75 m"},
		"2.5m * 0.0": {"2.5m", "0.0", "0 m"},
		// unlike
		"10S * 0m":          {"2.5m", "0m", "0 m^2"},
		"2.5m * 3N":         {"2.5m", "3N", "7.5 m*N"},
		"2.5 m^2 * 3kg/m^2": {"2.5 m^2", "3kg/m^2", "7.5 kg"},
		// inverse
		"10S * 2/S":   {"10S", "2/S", "20"},
		"2/S * 10S":   {"2/S", "10S", "20"},
		"10S * 0.1/S": {"10S", "0.1/S", "1"},
		"0.1/S * 10S": {"0.1/S", "10S", "1"},
		//
		"3m * 4 1/km":    {"3m", "4 1/km", "0.012"},
		"4 1/km * 3m":    {"4 1/km", "3m", "0.012"},
		"4 m * 3 A/km":   {"4m", "3 A/km", "0.012 A"},
		"3 A/km * 4 m":   {"3 A/km", "4m", "0.012 A"},
		"4 m * 3 1/km^2": {"4m", "3 1/km^2", "0.000012 1/m"},
		"3 1/km^2 * 4m":  {"3 1/km^2", "4m", "0.012 1/km"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := Parse(test.a)
			if err != nil {
				t.Errorf("failed to parse %v, got %v", test.a, err)
				return
			}
			b, err := Parse(test.b)
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

func TestDiv(t *testing.T) {
	tests := map[string]struct {
		a        string
		b        string
		expected string
	}{
		// degrees
		"7.5degF / 2.5m^2":  {"7.5degF", "2.5m^2", "3 \u00b0F/m^2"},
		"2.5 degF / 0 degF": {"2.5 degF", "0 degF", "divide by zero"},
		"2.5 degF / 0":      {"2.5 degF", "0", "divide by zero"},
		"2.5 degF / 4degF":  {"2.5 degF", "4degF", "0.625"},
		"2.5 degF / 4":      {"2.5 degF", "4", "0.625 \u00b0F"},
		"2.5 degF / 2 degC": {"2.5 degF", "2 degC", "1.25 \u00b0F/\u00b0C"},
		// temperature
		"tempF / 1 tempC": {"tempF", "1 tempC", "cannot divide with temperatures"},
		"tempF / 1 degC":  {"tempF", "1 degC", "cannot divide with temperatures"},
		"2 / tempF":       {"2", "tempF", "cannot divide with temperatures"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := Parse(test.a)
			if err != nil {
				t.Errorf("failed to parse %v, got %v", test.a, err)
				return
			}
			b, err := Parse(test.b)
			if err != nil {
				t.Errorf("failed to parse %v got %v", test.b, err)
				return
			}
			if actual, err := a.Div(b); err != nil {
				if err.Error() != test.expected {
					t.Errorf("expected %v, got %v", test.expected, err)
				}
			} else {
				str := actual.String()
				if str != test.expected {
					t.Errorf("expected %v, got %v", test.expected, actual)
				}
			}
			if actual, err := a.Div(test.b); err != nil {
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
