package brdocs

import (
	"fmt"
	"github.com/brazanation/go-documents/internal"
	"github.com/brazanation/go-documents/internal/calculator"
	"strconv"
)

const VoterType internal.DocumentType = "Voter"

const (
	voterLength         int    = 12
	voterNumberOfDigits int    = 2
	voterValidatorRegex string = `^([\d]{12})$`
	voterFormatterRegex string = "$1"
)

type voter struct {
}

// NewVoter is the constructor of voter
func NewVoter(number string) (internal.Document, error) {
	d, err := internal.NewDocument(
		VoterType,
		number,
		voterLength,
		voterNumberOfDigits,
		voterFormatterRegex,
		voterValidatorRegex,
		voter{},
	)
	return d, err
}

func (v voter) CalculateDigit(baseNumber string) string {
	firstDigit := v.calculateFirstDigit(baseNumber[:len(baseNumber)-2])
	bn := baseNumber+strconv.Itoa(firstDigit)
	secondDigit := v.calculateSecondDigit(bn[len(bn)-3:])
	return fmt.Sprintf("%d%d", firstDigit, secondDigit)
}

func (v voter) calculateFirstDigit(baseNumber string) int {
	m := calculator.NewModule11(baseNumber)
	m.WithMultipliers(9, 8, 7, 6, 5, 4, 3, 2)
	return m.Calculate()
}

func (v voter) calculateSecondDigit(baseNumber string) int {
	m := calculator.NewModule11(baseNumber)
	m.WithMultipliers(9, 8, 7)
	return m.Calculate()
}
