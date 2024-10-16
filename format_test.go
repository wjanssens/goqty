package goqty

import (
	"testing"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func TestDefaultFormatter(t *testing.T) {
	qty, err := ParseQty("2.987654321 m")
	if err != nil {
		t.Errorf("failed to create '2.987654321 m', got %v", err)
		return
	}
	f := qty.Format(DefaultFormatter)
	if err != nil {
		t.Errorf("failed to format, got %v", err)
		return
	}
	expected := "2.987654321 m"
	if f != expected {
		t.Errorf("expected formatted %v, got %v", expected, f)
	}
}

func TestFormatter(t *testing.T) {
	qty, err := ParseQty("2.987654321 m")
	if err != nil {
		t.Errorf("failed to create '2.987654321 m', got %v", err)
		return
	}
	fn := func(scalar float64, unit string) string {
		v := number.Decimal(scalar, number.MaxIntegerDigits(4), number.MinIntegerDigits(2))
		return message.NewPrinter(language.Dutch).Sprintf("%v %v", v, unit)
	}

	f := qty.Format(fn)
	if err != nil {
		t.Errorf("failed to format, got %v", err)
		return
	}
	expected := "02,988 m"
	if f != expected {
		t.Errorf("expected formatted %v, got %v", expected, f)
	}
}
