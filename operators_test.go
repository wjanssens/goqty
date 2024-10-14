package goqty

import "testing"

func TestAdd(t *testing.T) {
	qty1, err := ParseQty("2.5m")
	if err != nil {
		t.Errorf("failed to parse '2.5m', got %v", err)
		return
	}
	qty2, err := ParseQty("3m")
	if err != nil {
		t.Errorf("failed to parse '2.5m', got %v", err)
		return
	}

	if a, err := qty1.Add(qty2); err != nil {
		t.Errorf("failed to add, got %v", err)
	} else {
		if a.scalar != 5.5 {
			t.Errorf("expected scalar %v, got %v", 5.5, a.scalar)
		}
	}

	if a, err := qty1.Add("3m"); err != nil {
		t.Errorf("failed to add, got %v", err)
	} else {
		if a.scalar != 5.5 {
			t.Errorf("expected scalar %v, got %v", 5.5, a.scalar)
		}
	}

	qty2, err = ParseQty("3cm")
	if err != nil {
		t.Errorf("failed to parse '3cm', got %v", err)
		return
	}
	if a, err := qty1.Add(qty2); err != nil {
		t.Errorf("failed to add, got %v", err)
	} else {
		if a.scalar != 5.53 {
			t.Errorf("expected scalar %v, got %v", 5.53, a.scalar)
		}
	}
	if a, err := qty2.Add(qty1); err != nil {
		t.Errorf("failed to add, got %v", err)
	} else {
		if a.scalar != 253 {
			t.Errorf("expected scalar %v, got %v", 253, a.scalar)
		}
		u := a.Units()
		if u != "cm" {
			t.Errorf("expected units %v, got %v", "cm", u)
		}
	}

	qty1, err = ParseQty("5cm")
	if err != nil {
		t.Errorf("failed to parse '5cm', got %v", err)
		return
	}
	if a, err := qty2.Add(qty1); err != nil {
		t.Errorf("failed to add, got %v", err)
	} else {
		if a.scalar != 253 {
			t.Errorf("expected scalar %v, got %v", 253, a.scalar)
		}
		u := a.Units()
		if u != "cm" {
			t.Errorf("expected units %v, got %v", "cm", u)
		}
	}
}
