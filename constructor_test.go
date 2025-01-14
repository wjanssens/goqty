package goqty

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		scalar      float64
		units       string
		kind        string
		numerator   []string
		denominator []string
	}{
		"1.5":   {1.5, "", "unitless", []string{"<1>"}, []string{"<1>"}},
		"1.5m":  {1.5, "m", "length", []string{"<meter>"}, []string{"<1>"}},
		"-1.5m": {-1.5, "m", "length", []string{"<meter>"}, []string{"<1>"}},
		// compound
		"5 N*m": {5, "N*m", "energy", []string{"<newton>", "<meter>"}, []string{"<1>"}},
		"1 m/s": {1, "m/s", "speed", []string{"<meter>"}, []string{"<second>"}},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := New(test.scalar, test.units)
			if err != nil {
				t.Errorf("failed to create, got %v", err)
				return
			} else {
				if actual.scalar != test.scalar {
					t.Errorf("expected scalar %v, got %v", test.scalar, actual.scalar)
				}
				kind := actual.Kind()
				if kind != test.kind {
					t.Errorf("expected kind %v, got %v", test.kind, kind)
				}
			}
		})
	}
}
