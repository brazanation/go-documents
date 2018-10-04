package brdocument

import (
	"testing"
)

func TestShouldFormatPisPasepDocument(t *testing.T) {
	t.Run("PIS/PASEP 42585742143", func(t *testing.T) {
		pispasep := Document.NewPisPasep(Document{}, "42585742143")
		if pispasep.FormatPisPasep() != "425.85742.14-3" {
			t.Error("PIS/PASEP was not formatted")
		}
	})
}
