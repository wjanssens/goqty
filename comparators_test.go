package goqty

import "testing"

func TestComparisson(t *testing.T) {
	tests := map[string]struct {
		a        string
		op       string
		b        string
		expected bool
	}{
		"1 m2 == 1 m^2":                        {"1 m2", "eq", "1 m^2", true},
		"1 m^2 kg^2 J^2/s^2 == 1 m2 kg2 J2/s2": {"1 m^2 kg^2 J^2/s^2", "eq", "1 m2 kg2 J2/s2", true},
		"1cm == 10mm":                          {"1cm", "eq", "10mm", true},
		"1cm < 1mm":                            {"1cm", "lt", "1mm", false},
		"1cm < 10mm":                           {"1cm", "lt", "10mm", false},
		"1cm <= 1mm":                           {"1cm", "lte", "10mm", true},
		"1cm >= 1mm":                           {"1cm", "gte", "10mm", true},
		"1cm > 1mm":                            {"1cm", "gt", "1mm", true},
		"1mm > 1cm":                            {"1mm", "gt", "1cm", false},
		// same
		"1cm === 1cm":  {"1cm", "same", "1cm", true},
		"1cm !== 10mm": {"1cm", "same", "10mm", false},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := Parse(test.a)
			if err != nil {
				t.Errorf("failed to create %v, got %v", test.a, err)
				return
			}
			b, err := Parse(test.b)
			if err != nil {
				t.Errorf("failed to create %v, got %v", test.b, err)
				return
			}
			switch test.op {
			case "eq":
				if actual, err := a.Eq(b); err != nil {
					t.Errorf("expected %v == %v => %v, got %v", a, b, test.expected, err)
				} else if actual != test.expected {
					t.Errorf("expected %v == %v => %v", a, b, test.expected)
				}
			case "lt":
				if actual, err := a.Lt(b); err != nil {
					t.Errorf("expected %v == %v => %v, got %v", a, b, test.expected, err)
				} else if actual != test.expected {
					t.Errorf("expected %v < %v => %v", a, b, test.expected)
				}
			case "lte":
				if actual, err := a.Lte(b); err != nil {
					t.Errorf("expected %v == %v => %v, got %v", a, b, test.expected, err)
				} else if actual != test.expected {
					t.Errorf("expected %v <= %v => %v", a, b, test.expected)
				}
			case "gt":
				if actual, err := a.Gt(b); err != nil {
					t.Errorf("expected %v == %v => %v, got %v", a, b, test.expected, err)
				} else if actual != test.expected {
					t.Errorf("expected %v > %v => %v", a, b, test.expected)
				}
			case "gte":
				if actual, err := a.Gte(b); err != nil {
					t.Errorf("expected %v == %v => %v, got %v", a, b, test.expected, err)
				} else if actual != test.expected {
					t.Errorf("expected %v >= %v => %v", a, b, test.expected)
				}
			case "same":
				if actual, err := a.Same(b); err != nil {
					t.Errorf("expected %v == %v => %v, got %v", a, b, test.expected, err)
				} else if actual != test.expected {
					t.Errorf("expected %v === %v => %v", a, b, test.expected)
				}
			}
		})
	}
}

func TestCompareTo(t *testing.T) {
	tests := map[string]struct {
		a        string
		b        string
		expected int
	}{
		"1cm > 1mm":  {"1cm", "1mm", 1},
		"1mm < 1cm":  {"1mm", "1cm", -1},
		"1cm = 10mm": {"1cm", "10mm", 0},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := Parse(test.a)
			if err != nil {
				t.Errorf("failed to create %v, got %v", test.a, err)
				return
			}
			b, err := Parse(test.b)
			if err != nil {
				t.Errorf("failed to create %v, got %v", test.b, err)
				return
			}
			if actual, err := a.CompareTo(b); err != nil {
				t.Errorf("expected eq %v, got %v", test.expected, err)
			} else if actual != test.expected {
				t.Errorf("expected eq %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestCompareToError(t *testing.T) {
	tests := map[string]struct {
		a        string
		b        string
		expected string
	}{
		"1cm != 28A": {"1cm", "20A", "incompatible units: cm and A"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := Parse(test.a)
			if err != nil {
				t.Errorf("failed to create %v, got %v", test.a, err)
				return
			}
			b, err := Parse(test.b)
			if err != nil {
				t.Errorf("failed to create %v, got %v", test.b, err)
				return
			}
			if actual, err := a.CompareTo(b); err != nil {
				if err.Error() != test.expected {
					t.Errorf("expected eq %v, got %v", test.expected, err)
				}
			} else {
				t.Errorf("expected eq %v, got %v", test.expected, actual)
			}
		})
	}
}
