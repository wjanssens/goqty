package goqty

import "testing"

func TestEq(t *testing.T) {
	tests := map[string]struct {
		a        string
		b        string
		expected bool
	}{
		"1 m2 eq 1 m^2":                        {"1 m2", "1 m^2", true},
		"1 m^2 kg^2 J^2/s^2 eq 1 m2 kg2 J2/s2": {"1 m^2 kg^2 J^2/s^2", "1 m2 kg2 J2/s2", true},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := ParseQty(test.a)
			if err != nil {
				t.Errorf("failed to create %v, got %v", test.a, err)
				return
			}
			b, err := ParseQty(test.b)
			if err != nil {
				t.Errorf("failed to create %v, got %v", test.b, err)
				return
			}
			actual := a.Eq(b)
			if actual != test.expected {
				t.Errorf("expected eq %v, got %v", test.expected, actual)
			}
		})
	}
}
