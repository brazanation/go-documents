package internal

import (
	"github.com/brazanation/go-documents/internal/calculator"
	"regexp"
)

type DocumentType string

type Formattable interface {
	Format() string
}

// Document base structure
type Document struct {
	label          DocumentType
	number         string
	digit          string
	length         int
	numberOfDigits int
	regexFormatter string
	regexValidator string
	calculator     calculator.DigitCalculatable
}

func NewDocument(
	label DocumentType,
	number string,
	length int,
	numberOfDigits int,
	regexFormatter string,
	regexValidation string,
	calculator calculator.DigitCalculatable) (Document, error) {
	regex := regexp.MustCompile(`/\D/`)
	n := regex.ReplaceAllString(number, ``)

	d := Document{
		label:          label,
		number:         n,
		length:         length,
		numberOfDigits: numberOfDigits,
		regexFormatter: regexFormatter,
		regexValidator: regexValidation,
		calculator:     calculator,
	}

	d.digit = d.extractCheckerDigit()

	err := d.Assert()

	if err != nil {
		return Document{label: label, number: number}, err
	}

	return d, nil
}

func (d Document) Is (t DocumentType) bool {
	return d.label == t
}

// 	String returns a string with the Document number unformatted
func (d Document) String() string {
	return d.number
}

// Format returns a string with the Document number formatted
func (d Document) Format() string {
	regex := regexp.MustCompile(d.regexValidator)
	return regex.ReplaceAllString(d.number, d.regexFormatter)
}

// extractCheckerDigit returns the verificator digit of a Document
func (d Document) extractCheckerDigit() string {
	runes := []rune(d.number)
	digits := d.length - d.numberOfDigits
	return string(runes[digits:d.length])
}

// extractBaseNumber returns the base number of a Cpf
func (d Document) extractBaseNumber() string {
	baseDocument := ""
	for k, digit := range []byte(d.number) {
		if k <= (len(d.number) - int(d.numberOfDigits) - 1) {
			baseDocument += string(digit)
		}
	}
	return baseDocument
}

// Assert do the verifications to know if a Cpf is valid and not empty
func (d *Document) Assert() error {
	if d.number == "" {
		return NewEmptyErr(d.label)
	}

	if !d.isValid() {
		return NewInvalidErr(d)
	}

	return nil
}

func (d Document) isValid() bool {
	baseNumber := d.extractBaseNumber()
	regex := regexp.MustCompile("^[" + string(baseNumber[0]) + "]+$")
	isRepeated := regex.MatchString(baseNumber)

	if isRepeated {
		return false
	}

	digit := d.calculator.CalculateDigit(baseNumber)

	return d.IsSameDigit(digit)
}

func (d *Document) IsSameDigit(digit string) bool {
	return d.digit == digit
}
