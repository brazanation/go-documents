package calculator_test

import (
	"github.com/brazanation/go-documents/internal/calculator"
	"testing"
)

func TestCnsPermanentShouldBeValid(t *testing.T) {
	c := calculator.NewCnsCalculator("24291201846000")
	digit := c.CalculateDigit("24291201846000")
	if digit != "5" {
		t.Errorf("Failed on asserting that digit equals to 5, but actual %s", digit)
	}
}

func TestCnsTemporaryShouldBeValid(t *testing.T) {
	c := calculator.NewCnsCalculator("742912018460004")
	digit := c.CalculateDigit("74291201846000")
	if digit != "4" {
		t.Errorf("Failed on asserting that digit equals to 4, but actual %s", digit)
	}
}
