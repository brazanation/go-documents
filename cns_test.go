package brdocs_test

import (
	"github.com/brazanation/go-documents"
	"github.com/brazanation/go-documents/internal/test"
	"testing"
)

func TestCnsPermanentShouldBeValid(t *testing.T) {
	doc, err := brdocs.NewCns("242912018460005")
	internal_test.AssertValidDocument(t, doc, err)
}

func TestCnsTemporaryShouldBeValid(t *testing.T) {
	doc, err := brdocs.NewCns("742912018460004")
	internal_test.AssertValidDocument(t, doc, err)
}

func TestCnsShouldBeSameType(t *testing.T) {
	doc, _ := brdocs.NewCns("242912018460005")
	internal_test.AssertDocumentType(t, doc, brdocs.CnsType)
}

func TestCnsShouldBeFormatted(t *testing.T) {
	doc, err := brdocs.NewCns("242912018460005")
	internal_test.AssertValidDocument(t, doc, err)
	internal_test.AssertDocumentFormatted(t, doc, "242 9120 1846 0005")
}

func TestCnsShouldReturnErrorForRepeatedNumber(t *testing.T) {
	number := "111111111111111"
	_, err := brdocs.NewCns(number)
	internal_test.AssertNotValidDocument(t, number, err, "Cns(111111111111111) is invalid")
}

func TestCnsShouldReturnErrorForInvalidNumber(t *testing.T) {
	number := "861238979874098"
	_, err := brdocs.NewCns(number)
	internal_test.AssertNotValidDocument(t, number, err, "Cns(861238979874098) is invalid")
}

func TestCnsShouldReturnErrorForEmptyNumber(t *testing.T) {
	number := ``
	_, err := brdocs.NewCns(number)
	internal_test.AssertNotValidDocument(t, number, err, "Cns must not be empty")
}
