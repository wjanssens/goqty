package goqty

import "testing"

func TestTo(t *testing.T) {
	tests := map[string]struct {
		expr   string
		units  string
		scalar float64
		kind   string
	}{
		// // inverting kind
		"1 ohm to siemens":  {"1 ohm", "siemens", 1, "conductance"},
		"10 ohm to siemens": {"10 ohm", "siemens", 0.1, "conductance"},
		"10S to ohm":        {"10 S", "ohm", 0.1, "resistance"},
		"10 ohm to S":       {"10 ohm", "S", 0.1, "conductance"},
		// same kind
		"12 in to ft": {"12 in", "ft", 1, "length"},
		// same units
		"123 cm3":     {"123 cm3", "cm3", 123, "volume"},
		"123 cm3 alt": {"123 cm3", "cm^3", 123, "volume"},
		"123 ug":      {"123 ug", "µg", 123, "mass"},
		//
		"3 A/km": {"3 A/km", "3 A/m", 0.003, "magnetism"},
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
