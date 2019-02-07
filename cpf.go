package brdocs

import (
	"fmt"
	"github.com/brazanation/go-documents/internal"
	"github.com/brazanation/go-documents/internal/calculator"
)

const CpfType internal.DocumentType = "Cpf"

const (
	cpfLength         int    = 11
	cpfNumberOfDigits int    = 2
	cpfValidatorRegex string = `^([\d]{3})([\d]{3})([\d]{3})([\d]{2})$`
	cpfFormatterRegex string = "$1.$2.$3-$4"
)

type cpf struct {
}

// NewCpf is the constructor of cpf
func NewCpf(number string) (internal.Document, error) {
	d, err := internal.NewDocument(
		CpfType,
		number,
		cpfLength,
		cpfNumberOfDigits,
		cpfFormatterRegex,
		cpfValidatorRegex,
		cpf{},
	)
	return d, err
}

func (c cpf) CalculateDigit(base string) string {
	m := calculator.NewModule11(base)
	m.WithMultipliersInterval(cpfNumberOfDigits, cpfLength)
	m.UseComplementaryInsteadOfModule()
	m.ReplaceWhen(0, 10, 11)
	firstDigit := m.Calculate()
	m.AddDigit(firstDigit)
	secondDigit := m.Calculate()
	return fmt.Sprintf("%d%d", firstDigit, secondDigit)
}
