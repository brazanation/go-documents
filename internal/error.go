package internal

import "fmt"

type EmptyErr struct {
	label DocumentType
}

func (e EmptyErr) Error() string {
	return fmt.Sprintf("%s must not be empty", e.label)
}

func NewEmptyErr(label DocumentType) error {
	return EmptyErr{label: label}
}

type InvalidErr struct {
	label  DocumentType
	number string
}

func (e InvalidErr) Error() string {
	return fmt.Sprintf("%s(%s) is invalid", e.label, e.number)
}

func NewInvalidErr(doc *Document) error {
	return InvalidErr{label: doc.label, number: doc.number}
}
