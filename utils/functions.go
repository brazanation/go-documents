package utils

import (
	"regexp"
)

// ToString returns a string with the Document number unformatted
func (document Document) ToString() string {
	return document.Number
}

// Format returns a string with the Document number formatted
func (document Document) Format() string {
	regex := regexp.MustCompile(document.RegexValidator)
	return regex.ReplaceAllString(document.Number, document.RegexFormatter)
}

// ExtractCheckerDigit returns the verificator digit of a Document
func (document Document) ExtractCheckerDigit() string {
	runes := []rune(document.Number)
	digits := document.Length - document.NumberOfDigits
	return string(runes[digits:document.Length])
}

// ExtractBaseNumber returns the base number of a CPF
func (document Document) ExtractBaseNumber() string {
	baseDocument := ""
	for k, digit := range []byte(document.Number) {
		if k <= (len(document.Number) - int(document.NumberOfDigits) - 1) {
			baseDocument += string(digit)
		}
	}
	return baseDocument
}
