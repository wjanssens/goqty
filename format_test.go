package goqty

import "testing"

func TestDefaultFormatter(t *testing.T) {
	qty, err := ParseQty("2.987654321 m")
	if err != nil {
		t.Errorf("failed to create '2.987654321 m', got %v", err)
		return
	}
	f, err := qty.Format(nil)
	if err != nil {
		t.Errorf("failed to format, got %v", err)
		return
	}
	expected := "2.987654321 m"
	if f != expected {
		t.Errorf("expected formatted %v, got %v", expected, f)
	}
}
