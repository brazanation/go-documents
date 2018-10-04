package brdocument

import (
	"regexp"
)

const (
	PISPASEP_LENGTH = 11
    PISPASEP_REGEX = `^([\d]{3})([\d]{5})([\d]{2})([\d]{1})$`
    PISPASEP_NUMBER_OF_DIGITS = 1
)

func (document Document) NewPisPasep(number string) (Document) {
	regex := regexp.MustCompile(`/\D/`)
	pispasep := regex.ReplaceAllString(number, ``)

	return document.NewDocument(pispasep, PISPASEP_LENGTH, PISPASEP_NUMBER_OF_DIGITS)
}

func (document Document) FormatPisPasep() (string) {
	regex := regexp.MustCompile(PISPASEP_REGEX)
	return regex.ReplaceAllString(document.number, "$1.$2.$3-$4")
}