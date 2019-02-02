package internal_test

import (
	"github.com/brazanation/go-documents/internal"
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

func TestDocument(t *testing.T) {
	t.Run("Should create a dummy document similar a Cpf", func(t *testing.T) {
		digit := "73"
		d, err := internal.NewDocument(
			testDocument,
			"06843273173",
			11,
			2,
			formatterRegex,
			validatorRegex,
			CalculatorStub{digit},
		)

		if err != nil {
			t.Errorf("Failed asserting that not occurred error on create new document")
		}

		if !d.IsSameDigit(digit) {
			t.Errorf("Failed asserting that digit is %s", digit)
		}

		if d.String() != "06843273173" {
			t.Errorf("Failed asserting that number is equal to %s", "06843273173")
		}

		if d.Format() != "068.432.731-73" {
			t.Errorf("Failed asserting that number is formatted as %s", "068.432.731-73")
		}
	})

	t.Run("Should not create a dummy document", func(t *testing.T) {
		d, err := internal.NewDocument(
			testDocument,
			"06843273173",
			11,
			2,
			"",
			"",
			CalculatorStub{"00"},
		)

		if err == nil {
			t.Errorf("Failed asserting that occurred error on create new document")
		}

		if err.Error() != "Test(06843273173) is invalid" {
			t.Errorf("Failed asserting that error message \"%s\" for invalid document", err.Error())
		}

		if !d.Is(testDocument) {
			t.Errorf("Failed asserting that document is a %s", testDocument)
		}
	})

	t.Run("Should not create a empty document", func(t *testing.T) {
		_, err := internal.NewDocument(
			testDocument,
			"",
			11,
			2,
			"",
			"",
			CalculatorStub{"00"},
		)

		if err == nil {
			t.Errorf("Failed asserting that occurred error on create empty document")
		}

		if err.Error() != "Test must not be empty" {
			t.Errorf("Failed asserting that error message \"%s\" for empty document", err.Error())
		}
	})

	t.Run("Should not create document with repeated number", func(t *testing.T) {
		_, err := internal.NewDocument(
			testDocument,
			"11111111111",
			11,
			2,
			"",
			"",
			CalculatorStub{""},
		)

		if err == nil {
			t.Errorf("Failed asserting that occurred error on create empty document")
		}

		if err.Error() != "Test(11111111111) is invalid" {
			t.Errorf("Failed asserting that error message \"%s\" for empty document", err.Error())
		}
	})

}
