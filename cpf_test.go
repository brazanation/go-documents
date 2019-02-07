package brdocs_test

import (
	"github.com/brazanation/go-documents"
	"github.com/brazanation/go-documents/internal/test"
	"testing"
)

func TestCpfShouldBeValid(t *testing.T) {
	cpf, err := brdocs.NewCpf("06843273173")
	internal_test.AssertValidDocument(t, cpf, err)
}

func TestCpfShouldBeFormatted(t *testing.T) {
	cpf, err := brdocs.NewCpf("06843273173")
	internal_test.AssertValidDocument(t, cpf, err)
	internal_test.AssertDocumentFormatted(t, cpf, "068.432.731-73")
}

func TestCpfShouldReturnErrorForRepeatedNumber(t *testing.T) {
	number := "11111111111"
	_, err := brdocs.NewCpf(number)
	internal_test.AssertNotValidDocument(t, number, err, "Cpf(11111111111) is invalid")
}

func TestCpfShouldReturnErrorForInvalidNumber(t *testing.T) {
	number := "12223455677"
	_, err := brdocs.NewCpf(number)
	internal_test.AssertNotValidDocument(t, number, err, "Cpf(12223455677) is invalid")
}

func TestCpfShouldReturnErrorForEmptyNumber(t *testing.T) {
	number := ``
	_, err := brdocs.NewCpf(number)
	internal_test.AssertNotValidDocument(t, number, err, "Cpf must not be empty")
}
