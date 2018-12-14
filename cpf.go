package brdocs

import (
	"errors"
	"fmt"
	"github.com/brazanation/go-documents/utils"
	"regexp"
)

type Cpf struct {}

var cpf = utils.Document{}

// NewCpf is the constructor of Cpf
func NewCpf(number string) (utils.Document, error) {
	regex := regexp.MustCompile(`/\D/`)
	cpfNumber := regex.ReplaceAllString(number, ``)

	cpf = utils.Document{
		Number:         cpfNumber,
		Length:         utils.CpfLength,
		NumberOfDigits: utils.CpfNumberOfDigits,
		RegexFormatter: utils.CpfFormatterRegex,
		RegexValidator: utils.CpfValidatorRegex,
	}
	cpf.Digit = cpf.ExtractCheckerDigit()

	err := Cpf{}.Assert()
	if err != nil {
		return utils.Document{}, err
	}

	return cpf, nil
}

// Assert do the verifications to know if a Cpf is valid and not empty
func (document Cpf) Assert() error {
	if cpf.Number == "" {
		return errors.New("cpf number can not be empty")
	}

	if !document.isValid() {
		return errors.New("cpf is not valid")
	}

	return nil
}

func (document Cpf) isValid() bool {
	baseNumber := cpf.ExtractBaseNumber()
	regex := regexp.MustCompile("^[" + string(baseNumber[0]) + "]+$")
	isRepeated := regex.MatchString(baseNumber)

	if isRepeated {
		return false
	}

	digit := document.calculateDigit()

	return digit == cpf.Digit
}

func (document Cpf) calculateDigit() string {
	calculator := utils.NewDigitCalculator(cpf.ExtractBaseNumber())
	calculator.WithMultipliersInterval(int(cpf.NumberOfDigits), int(cpf.Length))
	calculator.UseComplementaryInsteadOfModule()
	calculator.ReplaceWhen(0, []int{10, 11})
	calculator.WithModule(utils.Module11)

	firstDigit := calculator.Calculate()
	calculator.AddDigit(firstDigit)
	secondDigit := calculator.Calculate()
	return fmt.Sprintf("%d%d", firstDigit, secondDigit)
}
