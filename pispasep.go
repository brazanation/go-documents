package brdocs

import (
	"fmt"
	"github.com/brazanation/go-documents/internal"
	"github.com/brazanation/go-documents/internal/calculator"
)

const PisPasepType internal.DocumentType = "PisPasep"

const (
	pisPasepLength         int    = 11
	pisPasepNumberOfDigits int    = 1
	pisPasepValidatorRegex string = `^([\d]{3})([\d]{5})([\d]{2})([\d]{1})$`
	pisPasepFormatterRegex string = "$1.$2.$3-$4"
)

type pisPasep struct {
}

// NewPisPasep is the constructor of pisPasep
func NewPisPasep(number string) (internal.Document, error) {
	d, err := internal.NewDocument(
		PisPasepType,
		number,
		pisPasepLength,
		pisPasepNumberOfDigits,
		pisPasepFormatterRegex,
		pisPasepValidatorRegex,
		pisPasep{},
	)
	return d, err
}

func (pisPasep pisPasep) CalculateDigit(base string) string {
	c := calculator.NewModule11(base)
	c.WithMultipliersInterval(2, 9);
	c.UseComplementaryInsteadOfModule();
	c.ReplaceWhen(0, 10, 11);
	digit := c.Calculate();
	return fmt.Sprintf("%d", digit)
}
