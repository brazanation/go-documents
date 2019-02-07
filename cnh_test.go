package brdocs_test

import (
	"github.com/brazanation/go-documents"
	"github.com/brazanation/go-documents/internal/test"
	"testing"
)

func TestCnhShouldBeValid(t *testing.T) {
	cnh, err := brdocs.NewCnh("83592802666")
	internal_test.AssertValidDocument(t, cnh, err)
}

func TestCnhShouldBeSameType(t *testing.T) {
	cnh, _ := brdocs.NewCnh("83592802666")
	internal_test.AssertDocumentType(t, cnh, brdocs.CnhType)
}

func TestCnhShouldBeFormatted(t *testing.T) {
	cnh, err := brdocs.NewCnh("83592802666")
	internal_test.AssertValidDocument(t, cnh, err)
	internal_test.AssertDocumentFormatted(t, cnh, "83592802666")
}

func TestCnhShouldReturnErrorForRepeatedNumber(t *testing.T) {
	number := "11111111111"
	_, err := brdocs.NewCnh(number)
	internal_test.AssertNotValidDocument(t, number, err, "Cnh(11111111111) is invalid")
}

func TestCnhShouldReturnErrorForInvalidNumber(t *testing.T) {
	number := "83592802655"
	_, err := brdocs.NewCnh(number)
	internal_test.AssertNotValidDocument(t, number, err, "Cnh(83592802655) is invalid")
}

func TestCnhShouldReturnErrorForEmptyNumber(t *testing.T) {
	number := ``
	_, err := brdocs.NewCnh(number)
	internal_test.AssertNotValidDocument(t, number, err, "Cnh must not be empty")
}
