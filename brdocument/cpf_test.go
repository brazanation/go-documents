package brdocument

import (
	"testing"
)

func TestShouldFormatDocument(t *testing.T) {
	t.Run("CPF 06843273173", func(t *testing.T) {
		cpf := Document.NewCpf(Document{}, "06843273173")
		if cpf.Format() != "068.432.731-73" {
			t.Error("CPF was not formatted")
		}
	})
}

func TestShouldThrowExceptionForEmptyData(t *testing.T) {
	t.Run("CPF ''", func(t *testing.T) {
		cpf := Document.NewCpf(Document{}, "")
		if cpf.ToString() != "" {
			t.Error("CPF cannot be empty")
		}
	})
}
