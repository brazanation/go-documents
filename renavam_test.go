package brdocs_test

import (
	"github.com/brazanation/go-documents"
	"github.com/brazanation/go-documents/internal/test"

	"testing"
)

func TestRenavamShouldBeValid(t *testing.T) {
	renavam, err := brdocs.NewRenavam("61855253306")
	internal_test.AssertValidDocument(t, renavam, err)
}

func TestRenavamShouldBeFormatted(t *testing.T) {
	renavam, err := brdocs.NewRenavam("73197444810")
	internal_test.AssertValidDocument(t, renavam, err)
	internal_test.AssertDocumentFormatted(t, renavam, "7319.744481-0")
}

func TestRenavamShouldReturnErrorForRepeatedNumber(t *testing.T) {
	number := "11111111111"
	_, err := brdocs.NewRenavam(number)
	internal_test.AssertNotValidDocument(t, number, err, "Renavam(11111111111) is invalid")
}

func TestRenavamShouldReturnErrorForInvalidNumber(t *testing.T) {
	number := "61855253307"
	_, err := brdocs.NewRenavam(number)
	internal_test.AssertNotValidDocument(t, number, err, "Renavam(61855253307) is invalid")
}

func TestRenavamShouldReturnErrorForEmptyNumber(t *testing.T) {
	number := ``
	_, err := brdocs.NewRenavam(number)
	internal_test.AssertNotValidDocument(t, number, err, "Renavam must not be empty")
}
