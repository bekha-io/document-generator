package errs

import "errors"

var (
	ErrDocumentTemplateDataInvalid = errors.New("ErrDocumentTemplateDataInvalid")
	ErrDocumentAlreadySigned = errors.New("ErrDocumentAlreadySigned")
)
