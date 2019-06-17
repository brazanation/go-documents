package brdocs

import (
	"fmt"
	"github.com/brazanation/go-documents/internal"
	"github.com/brazanation/go-documents/internal/calculator"
)

const RenavamType internal.DocumentType = "Renavam"

const (
	renavamLength         int    = 11
	renavamNumberOfDigits int    = 1
	renavamValidatorRegex string = `^([\d]{4})([\d]{6})([\d]{1})$`
	renavamFormatterRegex string = "$1.$2-$3"
)

type renavam struct {}

// NewRenavam is the constructor of renavam
func NewRenavam(number string) (internal.Document, error) {
	d, err := internal.NewDocument(
		RenavamType,
		number,
		renavamLength,
		renavamNumberOfDigits,
		renavamFormatterRegex,
		renavamValidatorRegex,
		renavam{},
	)
	return d, err
}

func (r renavam) CalculateDigit(base string) string {
	m := calculator.NewModule11(base)
	m.WithMultipliers(2, 3, 4, 5, 6, 7, 8, 9, 2, 3)
	m.ReplaceWhen(0, 10)
	m.MultiplySumBy(10)
	return fmt.Sprintf("%d", m.Calculate())
}
