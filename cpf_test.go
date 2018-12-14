package brdocs_test

import (
	"github.com/brazanation/go-documents"
	"testing"
)

func TestShouldFormatCpf(t *testing.T) {
	t.Run("CPF 06843273173", func(t *testing.T) {
		cpf, err := brdocs.NewCpf("06843273173")
		if err == nil && cpf.Format() != "068.432.731-73" {
			t.Error("CPF was not formatted")
		}
	})
}

func TestShouldThrowExceptionWithRepeatedNumbers(t *testing.T) {
	t.Run("CPF 11111111111", func(t *testing.T) {
		cpf, err := brdocs.NewCpf("11111111111")
		if err == nil {
			t.Error("CPF" + cpf.ToString() + " is repeated and not valid")
		}
	})
}

func TestShouldThrowExceptionWithInvalidCpf(t *testing.T) {
	t.Run("CPF 12223455677", func(t *testing.T) {
		cpf, err := brdocs.NewCpf("12223455677")
		if err == nil {
			t.Error("CPF" + cpf.ToString() + " is invalid")
		}
	})
}

func TestShouldThrowExceptionForEmptyCpf(t *testing.T) {
	t.Run("CPF ''", func(t *testing.T) {
		cpf, err := brdocs.NewCpf("")
		if err != nil && cpf.ToString() != "" {
			t.Error("CPF cannot be empty")
		}
	})
}
