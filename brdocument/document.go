package brdocument

import (
	"errors"
	"regexp"
)

type Formattable interface {
	Format(number string) (string)
}

type AbstractDocument interface {
	NewDocument(number string, length int32, numberOfDigits int32, documentType string) (Document)
	CreateFromString(document string) (Document)
	ToString() (Document)
	CalculateDigit(baseNumber string) (string)
	ExtractBaseNumber(number string) (string)
	ExtractCheckerDigit(number string) (string)
	Formattable
}

type Document struct {
	number string
	digit string
	length int32
	numberOfDigits int32
}

func (document Document) NewDocument(number string, length int32, numberOfDigits int32) (Document) {
	document = Document{
		number:number,
		length:length,
		numberOfDigits:numberOfDigits,

	}
	document.digit = document.ExtractCheckerDigit(number)
	return document
}

func (document Document) ExtractCheckerDigit(number string) (string) {
	if document.numberOfDigits < 2 {
		return string(number[int32(len(number))-document.numberOfDigits])
	}
	runes := []rune(number)
	digits := document.length - document.numberOfDigits
	return string(runes[digits:document.length])
}

func (document Document) AssertDocument(number string) (error) {
	if number == "" {
		return errors.New("document number can not be empty")
	}

	if !document.isValid(number) {
		return errors.New("document is not valid")
	}

	return nil
}

func (document Document) isValid(number string) (bool) {
	baseNumber := document.ExtractBaseNumber(number)
	if baseNumber == "" {
		return false
	}

	regex := regexp.MustCompile("^[" + string(baseNumber[0]) + "]+$")
	isRepeated := regex.MatchString(baseNumber)

	if isRepeated {
		return false
	}

	digit := document.CalculateDigit()

	return digit == document.digit
}

func (document Document) ExtractBaseNumber(number string) (string) {
	baseDocument := ""
	for k, digit := range []byte(number) {
		if k <= (len(number) - int(document.numberOfDigits) - 1) {
			baseDocument += string(digit)
		}
	}
	return baseDocument
}