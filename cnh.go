package brdocs

import (
	"fmt"
	"github.com/brazanation/go-documents/internal"
	"github.com/brazanation/go-documents/internal/calculator"
)

const CnhType internal.DocumentType = "cnh"

const (
	cnhLength         int    = 11
	cnhNumberOfDigits int    = 2
	cnhValidatorRegex string = `^([\d]{11})$`
	cnhFormatterRegex string = "$1"
)

type cnh struct {
}

// NewCnh is the constructor of cnh
func NewCnh(number string) (internal.Document, error) {
	d, err := internal.NewDocument(
		CnhType,
		number,
		cnhLength,
		cnhNumberOfDigits,
		cnhFormatterRegex,
		cnhValidatorRegex,
		cnh{},
	)
	return d, err
}

func (c cnh) CalculateDigit(baseNumber string) string {
	firstDigit := c.calculateFirstDigit(baseNumber)
	secondDigit := c.calculateSecondDigit(baseNumber)
	return fmt.Sprintf("%d%d", firstDigit, secondDigit)
}

func (c cnh) calculateFirstDigit(baseNumber string) int {
	m := calculator.NewModule11(baseNumber)
	m.WithMultipliersInterval(1, 9)
	m.ReplaceWhen(0, 10, 11)
	return m.Calculate()
}

func (c cnh) calculateSecondDigit(baseNumber string) int {
	m := calculator.NewModule11(baseNumber)
	m.WithMultipliers(9, 8, 7, 6, 5, 4, 3, 2, 1)
	m.ReplaceWhen(0, 10, 11)
	return m.Calculate()
}
