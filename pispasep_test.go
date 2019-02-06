package brdocs_test

import (
	"github.com/brazanation/go-documents"
	"github.com/brazanation/go-documents/internal/test"
	"testing"
)

func TestPisPasepShouldBeValid(t *testing.T) {
	pis, err := brdocs.NewPisPasep("51823129491")
	internal_test.AssertValidDocument(t, pis, err)
}

func TestPisPasepShouldBeSameType(t *testing.T) {
	pis, _ := brdocs.NewPisPasep("51823129491")
	internal_test.AssertDocumentType(t, pis, brdocs.PisPasepType)
}

func TestPisPasepShouldBeFormatted(t *testing.T) {
	pis, err := brdocs.NewPisPasep("51823129491")
	internal_test.AssertValidDocument(t, pis, err)
	internal_test.AssertDocumentFormatted(t, pis, "518.23129.49-1")
}

func TestPisPasepShouldReturnErrorForRepeatedNumber(t *testing.T) {
	number := "11111111111"
	_, err := brdocs.NewPisPasep(number)
	internal_test.AssertNotValidDocument(t, number, err, "PisPasep(11111111111) is invalid")
}

func TestPisPasepShouldReturnErrorForInvalidNumber(t *testing.T) {
	number := "51823129490"
	_, err := brdocs.NewPisPasep(number)
	internal_test.AssertNotValidDocument(t, number, err, "PisPasep(51823129490) is invalid")
}

func TestPisPasepShouldReturnErrorForEmptyNumber(t *testing.T) {
	number := ``
	_, err := brdocs.NewPisPasep(number)
	internal_test.AssertNotValidDocument(t, number, err, "PisPasep must not be empty")
}
