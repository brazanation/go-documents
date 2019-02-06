package brdocs_test

import (
	"github.com/brazanation/go-documents"
	"github.com/brazanation/go-documents/internal/test"
	"testing"
)

func TestCnpjShouldBeValid(t *testing.T) {
	cnpj, err := brdocs.NewCnpj("99999090910270")
	internal_test.AssertValidDocument(t, cnpj, err)
}

func TestCnpjShouldBeSameType(t *testing.T) {
	cnpj, _ := brdocs.NewCnpj("99999090910270")
	internal_test.AssertDocumentType(t, cnpj, brdocs.CnpjType)
}

func TestCnpjShouldBeFormatted(t *testing.T) {
	cnpj, err := brdocs.NewCnpj("99999090910270")
	internal_test.AssertValidDocument(t, cnpj, err)
	internal_test.AssertDocumentFormatted(t, cnpj, "99.999.090/9102-70")
}

func TestCnpjShouldReturnErrorForRepeatedNumber(t *testing.T) {
	number := "11111111111111"
	_, err := brdocs.NewCnpj(number)
	internal_test.AssertNotValidDocument(t, number, err, "Cnpj(11111111111111) is invalid")
}

func TestCnpjShouldReturnErrorForInvalidNumber(t *testing.T) {
	number := "00111222100099"
	_, err := brdocs.NewCnpj(number)
	internal_test.AssertNotValidDocument(t, number, err, "Cnpj(00111222100099) is invalid")
}

func TestCnpjShouldReturnErrorForEmptyNumber(t *testing.T) {
	number := ``
	_, err := brdocs.NewCnpj(number)
	internal_test.AssertNotValidDocument(t, number, err, "Cnpj must not be empty")
}
