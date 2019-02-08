package brdocs

import (
	"github.com/brazanation/go-documents/internal"
	"github.com/brazanation/go-documents/internal/calculator"
)

const CnsType internal.DocumentType = "Cns"

const (
	cnsLength         int    = 15
	cnsNumberOfDigits int    = 1
	cnsValidatorRegex string = `^([\d]{3})([\d]{4})([\d]{4})([\d]{4})$`
	cnsFormatterRegex string = "$1 $2 $3 $4"
)

type cns struct {
}

// NewCns is the constructor of cns
func NewCns(number string) (internal.Document, error) {
	d, err := internal.NewDocument(
		CnsType,
		number,
		cnsLength,
		cnsNumberOfDigits,
		cnsFormatterRegex,
		cnsValidatorRegex,
		cns{},
	)
	return d, err
}

func (d cns) CalculateDigit(base string) string {
	c := calculator.NewCnsCalculator(base)
	digit := c.CalculateDigit(base)
	return digit
}
