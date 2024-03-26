package errs

import "errors"

var (
	ErrTemplateNotValid = errors.New("ErrTemplateNotValid")
	ErrTemplateNotFound = errors.New("ErrTemplateNotFound")
)
