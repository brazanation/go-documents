package brdocs

import (
	"fmt"
	"github.com/brazanation/go-documents/internal"
	"github.com/brazanation/go-documents/internal/calculator"
)

const CnpjType internal.DocumentType = "Cnpj"

const (
	cnpjLength         int    = 14
	cnpjNumberOfDigits int    = 2
	cnpjValidatorRegex string = `^([\d]{2})([\d]{3})([\d]{3})([\d]{4})([\d]{2})$`
	cnpjFormatterRegex string = "$1.$2.$3/$4-$5"
)

type Cnpj struct {
}

// NewCnpj is the constructor of Cnpj
func NewCnpj(number string) (internal.Document, error) {
	d, err := internal.NewDocument(
		CnpjType,
		number,
		cnpjLength,
		cnpjNumberOfDigits,
		cnpjFormatterRegex,
		cnpjValidatorRegex,
		Cnpj{},
	)
	return d, err
}

func (cnpj Cnpj) CalculateDigit(base string) string {
	c := calculator.NewModule11(base)
	c.UseComplementaryInsteadOfModule()
	c.ReplaceWhen(0, 10, 11)
	firstDigit := c.Calculate()
	c.AddDigit(firstDigit)
	secondDigit := c.Calculate()
	return fmt.Sprintf("%d%d", firstDigit, secondDigit)
}
