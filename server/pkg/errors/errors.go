package errors

import "errors"

var (
	ErrNotFound     = errors.New("resource not found")
	ErrAlreadyExist = errors.New("resource already exists")
	ErrInvalidInput = errors.New("invalid input")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrUpdateFailed = errors.New("failed to update resource")
	ErrCreateFailed = errors.New("failed to create resource")
	ErrDeleteFailed = errors.New("failed to delete resource")
)
