package brdocs_test

import (
	"github.com/brazanation/go-documents"
	"github.com/brazanation/go-documents/internal/test"
	"testing"
)

func TestVoterShouldBeValid(t *testing.T) {
	voter, err := brdocs.NewVoter("247003181023")
	internal_test.AssertValidDocument(t, voter, err)
}

func TestVoterShouldBeSameType(t *testing.T) {
	voter, _ := brdocs.NewVoter("247003181023")
	internal_test.AssertDocumentType(t, voter, brdocs.VoterType)
}

func TestVoterShouldBeFormatted(t *testing.T) {
	voter, err := brdocs.NewVoter("247003181023")
	internal_test.AssertValidDocument(t, voter, err)
	internal_test.AssertDocumentFormatted(t, voter, "247003181023")
}

func TestVoterShouldReturnErrorForRepeatedNumber(t *testing.T) {
	number := "1111111111111"
	_, err := brdocs.NewVoter(number)
	internal_test.AssertNotValidDocument(t, number, err, "Voter(1111111111111) is invalid")
}

func TestVoterShouldReturnErrorForInvalidNumber(t *testing.T) {
	number := "106644440301"
	_, err := brdocs.NewVoter(number)
	internal_test.AssertNotValidDocument(t, number, err, "Voter(106644440301) is invalid")
}

func TestVoterShouldReturnErrorForEmptyNumber(t *testing.T) {
	number := ``
	_, err := brdocs.NewVoter(number)
	internal_test.AssertNotValidDocument(t, number, err, "Voter must not be empty")
}
