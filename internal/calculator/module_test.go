package calculator_test

import (
	"fmt"
	"github.com/brazanation/go-documents/internal/calculator"
	"testing"
)

func TestDigitCalculator(t *testing.T) {
	t.Run("CPF 068432731-73", func(t *testing.T) {
		dc := calculator.NewModule11("068432731")
		dc.WithMultipliersInterval(2, 11)
		dc.UseComplementaryInsteadOfModule()
		dc.ReplaceWhen(0, 10, 11)
		d1 := dc.Calculate()
		dc.AddDigit(d1)
		d2 := dc.Calculate()

		digit := fmt.Sprintf("%d%d", d1, d2)
		expected := "73"

		if expected != digit {
			t.Errorf("Failed on asserting digit module 11 expecting %s and actual is %s", expected, digit)
		}
	})

	t.Run("Module 10 test", func(t *testing.T) {
		dc := calculator.NewModule10("1000003")
		dc.UseComplementaryInsteadOfModule()
		dc.ReplaceWhen(0, 10, 11)
		d1 := dc.Calculate()
		dc.AddDigit(d1)
		d2 := dc.Calculate()

		digit := fmt.Sprintf("%d%d", d2, d1)
		expected := "06"

		if digit != expected {
			t.Errorf("Failed on asserting digit for module 10 expecting %s and actual is %s", expected, digit)
		}
	})

	t.Run("calculate without using additional", func(t *testing.T) {
		dc1 := calculator.NewModule11("827894041")
		dc1.WithMultipliersInterval(1, 9)
		dc1.ReplaceWhen(0, 10, 11)
		first := dc1.Calculate()

		dc2 := calculator.NewModule11("827894041")
		dc2.WithMultipliers(9, 8, 7, 6, 5, 4, 3, 2, 1)
		dc2.ReplaceWhen(0, 10, 11)
		second := dc2.Calculate()

		digit := fmt.Sprintf("%d%d", first, second)
		expected := "20"

		if digit != expected {
			t.Errorf(
				"Failed on asserting digit for without additional expecting %s and actual is %s",
				expected,
				digit,
			)
		}
	})
}
