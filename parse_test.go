package qty

import (
	"slices"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		expr        string
		expected    string
		scalar      float64
		kind        string
		numerator   []string
		denominator []string
	}{
		// unit only
		"m": {"m", "1 m", 1, "length", []string{"<meter>"}, []string{"<1>"}},
		// unitless
		"1":   {"1", "1", 1, "unitless", []string{"<1>"}, []string{"<1>"}},
		"-1":  {"-1", "-1", -1, "unitless", []string{"<1>"}, []string{"<1>"}},
		"1.5": {"1.5", "1.5", 1.5, "unitless", []string{"<1>"}, []string{"<1>"}},
		// simple
		"1.5m":  {"1.5m", "1.5 m", 1.5, "length", []string{"<meter>"}, []string{"<1>"}},
		"-1.5m": {"-1.5m", "-1.5 m", -1.5, "length", []string{"<meter>"}, []string{"<1>"}},
		// denominator
		"1.5 /m": {"1.5 /m", "1.5 1/m", 1.5, "wavenumber", []string{"<1>"}, []string{"<meter>"}},
		"-1.5/m": {"-1.5 /m", "-1.5 1/m", -1.5, "wavenumber", []string{"<1>"}, []string{"<meter>"}},
		"5 1/m":  {"5 1/m", "5 1/m", 5, "wavenumber", []string{"<1>"}, []string{"<meter>"}},
		// prefixes including unicode prefixes
		"1 um":      {"1 um", "1 \u00B5m", 1, "length", []string{"<micro>", "<meter>"}, []string{"<1>"}},
		"1 \u03BCm": {"1 \u03BCm", "1 \u00B5m", 1, "length", []string{"<micro>", "<meter>"}, []string{"<1>"}}, // Greek
		"1 \u00B5m": {"1 \u00B5m", "1 \u00B5m", 1, "length", []string{"<micro>", "<meter>"}, []string{"<1>"}}, // micro
		"1 ohm":     {"1 ohm", "1 \u2126", 1, "resistance", []string{"<ohm>"}, []string{"<1>"}},
		"1 \u03A9":  {"1 \u03A9", "1 \u2126", 1, "resistance", []string{"<ohm>"}, []string{"<1>"}}, // Greek
		"1 \u2126":  {"1 \u2126", "1 \u2126", 1, "resistance", []string{"<ohm>"}, []string{"<1>"}}, // ohm
		"1 \u00b0":  {"1 \u00b0", "1 \u00b0", 1, "angle", []string{"<degree>"}, []string{"<1>"}},
		"1 \u00b0C": {"1 \u00b0C", "1 \u00b0C", 1, "temperature", []string{"<celsius>"}, []string{"<1>"}},
		// compound
		"5 N*m":  {"5 N*m", "5 N*m", 5, "energy", []string{"<newton>", "<meter>"}, []string{"<1>"}},
		"3 A/km": {"3 A/km", "3 A/km", 3, "magnetism", []string{"<ampere>"}, []string{"<kilo>", "<meter>"}},
		"1 m/s":  {"1 m/s", "1 m/s", 1, "speed", []string{"<meter>"}, []string{"<second>"}},
		// pressure (negative lookahead)
		"1 inH2O": {"1 inH2O", "1 inH2O", 1, "pressure", []string{"<inh2o>"}, []string{"<1>"}},
		"1 cmH2O": {"1 cmH2O", "1 cmH2O", 1, "pressure", []string{"<cmh2o>"}, []string{"<1>"}},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			q, err := Parse(test.expr)
			if err != nil {
				t.Errorf("failed to parse %v, got %v", test.expr, err)
			} else {
				str := q.String()
				if str != test.expected {
					t.Errorf("expected string %v, got %v", test.expected, str)
				}
				if q.scalar != test.scalar {
					t.Errorf("expected scalar %v, got %v", test.scalar, q.scalar)
				}
				kind := q.Kind()
				if kind != test.kind {
					t.Errorf("expected kind %v, got %v", test.kind, kind)
				}
				if !slices.Equal(q.numerator, test.numerator) {
					t.Errorf("expected numerator %v, got %v", test.numerator, q.numerator)
				}
				if !slices.Equal(q.denominator, test.denominator) {
					t.Errorf("expected denominator %v, got %v", test.denominator, q.denominator)
				}
			}
		})
	}
}

func TestParseFailure(t *testing.T) {
	tests := map[string]struct {
		expr     string
		expected string
	}{
		// temperature division
		// "2 tempF/s":               {"2 tempF/s", "cannot divide with temperatures"},
		// "2 s/tempF":               {"2 2 s/tempF", "cannot divide with temperatures"},
		"593720475cm^4939207503":  {"593720475cm^4939207503", "unit exponent is not a number"},
		"593720475cm**4939207503": {"593720475cm**4939207503", "unit exponent is not a number"},
		"593720475cm^5":           {"593720475cm^5", "unit not recognized"},
		"593720475cm**55":         {"593720475cm**55", "unit not recognized"},
		"aa":                      {"aa", "unit not recognized"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			q, err := Parse(test.expr)
			if err != nil {
				if err.Error() != test.expected {
					t.Errorf("expected error %v, got %v", test.expected, err)
				}
			} else {
				t.Errorf("expected error %v, got %v", test.expected, q)
			}
		})
	}
}
