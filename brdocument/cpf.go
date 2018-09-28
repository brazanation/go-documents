package brdocument

import (
	"regexp"
	"fmt"
)

const (
	CPF_LENGTH = 11
	CPF_REGEX = `^([\d]{3})([\d]{3})([\d]{3})([\d]{2})$`
	CPF_NUMBER_OF_DIGITS = 2
)

func (document Document) NewCpf(number string) (Document) {
	regex := regexp.MustCompile(`/\D/`)
	cpf := regex.ReplaceAllString(number, ``)

	return document.NewDocument(cpf, CPF_LENGTH, CPF_NUMBER_OF_DIGITS)
}

func (document Document) ToString() (string) {
	return document.number
}

func (document Document) Format() (string) {
	regex := regexp.MustCompile(CPF_REGEX)
	return regex.ReplaceAllString(document.number, "$1.$2.$3-$4")
}

func (document Document) CalculateDigit() (string) {
	calculator := NewDigitCalculator(document.ExtractBaseNumber(document.number))
	calculator.WithMultipliersInterval(CPF_NUMBER_OF_DIGITS, CPF_LENGTH)
	calculator.UseComplementaryInsteadOfModule()
	calculator.ReplaceWhen(0, []int{10,11})
	calculator.WithModule(Module11)

	firstDigit := calculator.Calculate()
	calculator.AddDigit(firstDigit)
	secondDigit := calculator.Calculate()
	return fmt.Sprintf("%d%d", firstDigit, secondDigit)
}
