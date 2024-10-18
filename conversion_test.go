package goqty

import "testing"

func TestTo(t *testing.T) {
	tests := map[string]struct {
		expr   string
		units  string
		scalar float64
		kind   string
	}{
		// inverting kind
		"1 ohm -> siemens":  {"1 ohm", "siemens", 1, "conductance"},
		"10 ohm -> siemens": {"10 ohm", "siemens", 0.1, "conductance"},
		"10S -> ohm":        {"10 S", "ohm", 0.1, "resistance"},
		"10 ohm -> S":       {"10 ohm", "S", 0.1, "conductance"},
		// same kind
		"12 in -> ft": {"12 in", "ft", 1, "length"},
		// same units
		"123 cm3 -> cm3":  {"123 cm3", "cm3", 123, "volume"},
		"123 cm3 -> cm^3": {"123 cm3", "cm^3", 123, "volume"},
		"123 ug -> µg":    {"123 ug", "µg", 123, "mass"},
		//
		"3 A/km -> 3 A/m": {"3 A/km", "3 A/m", 0.003, "magnetism"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			qty, err := ParseQty(test.expr)
			if err != nil {
				t.Errorf("failed to parse %v, got %v", test.expr, err)
				return
			}
			if actual, err := qty.To(test.units); err != nil {
				if err.Error() != test.kind {
					t.Errorf("expected %v, got %v", test.kind, err)
				}
			} else {
				if actual.scalar != test.scalar {
					t.Errorf("expected scalar %v, got %v", test.scalar, actual.scalar)
				}
				kind := actual.Kind()
				if kind != test.kind {
					t.Errorf("expected kind %v, got %v", test.scalar, kind)
				}
			}
		})
	}
}

func TestToPrec(t *testing.T) {
	tests := map[string]struct {
		q        string
		p        string
		expected string
	}{
		"5.17ft to ft => 5 ft":         {"5.17ft", "ft", "5 ft"},
		"5.17ft to 2ft => 6 ft":        {"5.17ft", "2ft", "6 ft"},
		"5.17ft to 10ft => 10 ft":      {"5.17ft", "10ft", "10 ft"},
		"5.17ft to 0.5ft => 5 ft":      {"5.17ft", "0.5ft", "5 ft"},
		"5.17ft to 0.25ft => 5.17 ft":  {"5.17ft", "0.25ft", "5.25 ft"},
		"5.17ft to 0.1ft => 5.2 ft":    {"5.17ft", "0.1ft", "5.2 ft"},
		"5.17ft to 0.05ft => 5.15 ft":  {"5.17ft", "0.05ft", "5.15 ft"},
		"5.17ft to 0.01ft => 5.17 ft":  {"5.17ft", "0.01ft", "5.17 ft"},
		"5.17ft to 0.001ft => 5.17 ft": {"5.17ft", "0.001ft", "5.17 ft"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			qty, err := ParseQty(test.q)
			if err != nil {
				t.Errorf("failed to parse %v, got %v", test.q, err)
				return
			}
			if actual, err := qty.ToPrec(test.p); err != nil {
				if err.Error() != test.expected {
					t.Errorf("expected %v, got %v", test.expected, err)
				}
			} else {
				string := actual.String()
				if string != test.expected {
					t.Errorf("expected %v, got %v", test.expected, string)
				}
			}
		})
	}
}
