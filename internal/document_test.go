package internal_test

import (
	"github.com/brazanation/go-documents/internal"
	"github.com/brazanation/go-documents/internal/test"
	"testing"
)

const (
	testDocument   internal.DocumentType = "Test"
	validatorRegex string                = `^([\d]{3})([\d]{3})([\d]{3})([\d]{2})$`
	formatterRegex string                = "$1.$2.$3-$4"
)

type CalculatorStub struct {
	result string
}

func (c CalculatorStub) CalculateDigit(base string) string {
	return c.result
}

func TestDocumentShouldBeValid(t *testing.T) {
	d, err := internal.NewDocument(
		testDocument,
		"06843273173",
		11,
		2,
		formatterRegex,
		validatorRegex,
		CalculatorStub{"73"},
	)
	internal_test.AssertValidDocument(t, d, err)
}

func TestDocumentShouldBeSameType(t *testing.T) {
	d, _ := internal.NewDocument(
		testDocument,
		"06843273173",
		11,
		2,
		formatterRegex,
		validatorRegex,
		CalculatorStub{"73"},
	)
	internal_test.AssertDocumentType(t, d, testDocument)
}

func TestDocumentShouldBeFormatted(t *testing.T) {
	d, err := internal.NewDocument(
		testDocument,
		"06843273173",
		11,
		2,
		formatterRegex,
		validatorRegex,
		CalculatorStub{"73"},
	)
	internal_test.AssertValidDocument(t, d, err)
	internal_test.AssertDocumentFormatted(t, d, "068.432.731-73")
}

func TestDocumentShouldReturnErrorForRepeatedNumber(t *testing.T) {
	number := "11111111111111"
	_, err := internal.NewDocument(
		testDocument,
		number,
		11,
		2,
		formatterRegex,
		validatorRegex,
		CalculatorStub{""},
	)
	internal_test.AssertNotValidDocument(t, number, err, "Test(11111111111111) is invalid")
}

func TestDocumentShouldReturnErrorForInvalidNumber(t *testing.T) {
	number := "00111222100099"
	_, err := internal.NewDocument(
		testDocument,
		number,
		11,
		2,
		formatterRegex,
		validatorRegex,
		CalculatorStub{""},
	)
	internal_test.AssertNotValidDocument(t, number, err, "Test(00111222100099) is invalid")
}

func TestDocumentShouldReturnErrorForEmptyNumber(t *testing.T) {
	number := ``
	_, err := internal.NewDocument(
		testDocument,
		number,
		11,
		2,
		formatterRegex,
		validatorRegex,
		CalculatorStub{""},
	)
	internal_test.AssertNotValidDocument(t, number, err, "Test must not be empty")
}
