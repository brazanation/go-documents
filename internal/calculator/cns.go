package calculator

import "fmt"

type cnsPermanent struct{}

func (p cnsPermanent) CalculateDigit(baseNumber string) string {
	n := baseNumber[:11]
	m := NewModule11(n)
	m.UseComplementaryInsteadOfModule()
	m.ReplaceWhen(8, 10)
	m.ReplaceWhen(0, 11)
	m.WithMultipliersInterval(5, 15)
	digit := m.Calculate()
	return fmt.Sprintf("%d", digit)
}

type cnsTemporary struct{}

func (t cnsTemporary) CalculateDigit(baseNumber string) string {
	m := NewModule11(baseNumber)
	m.WithMultipliersInterval(5, 15)
	digit := m.Calculate()
	return fmt.Sprintf("%d", digit)
}

func NewCnsCalculator(baseNumber string) DigitCalculatable {
	if isCnsTemporaryNumber(baseNumber) {
		return cnsTemporary{}
	}
	return cnsPermanent{}
}

func isCnsTemporaryNumber(baseNumber string) bool {
	number := baseNumber[:1]
	numbers := []string{"7", "8", "9"}
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}
