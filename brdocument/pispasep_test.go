package brdocument

import (
	"testing"
)

func TestShouldFormatPisPasepDocument(t *testing.T) {
	t.Run("PIS/PASEP 42585742141", func(t *testing.T) {
		pispasep := Document.NewPisPasep(Document{}, "42585742141")
		if pispasep.FormatPisPasep() != "425.85742.14-1" {
			t.Error("PIS/PASEP was not formatted")
		}
	})
}

func TestCalculatePisPasepDigit(t *testing.T) {
	t.Run("PIS/PASEP 42585742141", func(t *testing.T) {
		pispasep := Document.NewPisPasep(Document{}, "42585742141")
		if pispasep.CalculatePisPasepDigit() != "1" {
			t.Error("PIS/PASEP digit are != of 1")
		}
	})
}
