package internal_test

import (
	"fmt"
	"github.com/brazanation/go-documents/internal"
	"testing"
)

func AssertValidDocument(t *testing.T, d internal.Document, err error) {
	t.Run(fmt.Sprintf("assert valid Document(%s)", d.String()), func(t *testing.T) {
		if err != nil {
			t.Errorf("Failed on asserting none error occurred: %s", err.Error())
		}
	})
}

func AssertDocumentType(t *testing.T, d internal.Document, dt internal.DocumentType) {
	t.Run(fmt.Sprintf("assert Document Type(%s)", dt), func(t *testing.T) {
		if !d.Is(dt) {
			t.Errorf("Failed on asserting document type is %s", dt)
		}
	})
}

func AssertDocumentFormatted(t *testing.T, d internal.Document, formatted string) {
	t.Run(fmt.Sprintf("assert valid Document(%s)", d.String()), func(t *testing.T) {
		f := d.Format()
		if f != formatted {
			t.Errorf("Failed on asserting that document format expected is %s, but actual %s", formatted, f)
		}
	})
}

func AssertNotValidDocument(t *testing.T, number string, err error, expected string) {
	t.Run(fmt.Sprintf("assert valid Document(%s)", number), func(t *testing.T) {
		if err != nil && err.Error() != expected {
			t.Errorf("Failed on asserting invalid error message: %s", err.Error())
		}
	})
}
