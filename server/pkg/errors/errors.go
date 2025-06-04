package errors

import (
	"errors"
	"fmt"
)

// default errors
var (
	ErrInvalidInput   = errors.New("invalid input")           // 400
	ErrUnauthorized   = errors.New("unauthorized")            // 401
	ErrForbidden      = errors.New("forbidden")               // 403
	ErrNotFound       = errors.New("resource not found")      // 404
	ErrAlreadyExist   = errors.New("resource already exists") // 409
	ErrConflict       = errors.New("conflict")                // 409
	ErrTooManyRequest = errors.New("too many requests")       // 429

	ErrInternalServer  = errors.New("internal server error")     // 500
	ErrCreateFailed    = errors.New("failed to create resource") // 500
	ErrUpdateFailed    = errors.New("failed to update resource") // 500
	ErrDeleteFailed    = errors.New("failed to delete resource") // 500
	ErrTokenGeneration = errors.New("failed to generate token")  // 500
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

// customize message use for fallback
func NewBadRequest(message string) *AppError {
	return &AppError{Code: 400, Message: message, Err: ErrInvalidInput}
}

func NewUnauthorized(message string) *AppError {
	return &AppError{Code: 401, Message: message, Err: ErrUnauthorized}
}

func NewForbidden(message string) *AppError {
	return &AppError{Code: 403, Message: message, Err: ErrForbidden}
}

func NewNotFound(message string) *AppError {
	return &AppError{Code: 404, Message: message, Err: ErrNotFound}
}

func NewConflict(message string) *AppError {
	return &AppError{Code: 409, Message: message, Err: ErrConflict}
}

func NewAlreadyExist(message string) *AppError {
	return &AppError{Code: 409, Message: message, Err: ErrAlreadyExist}
}

func NewTooManyRequest(message string) *AppError {
	return &AppError{Code: 429, Message: message, Err: ErrTooManyRequest}
}

func NewInternal(message string, err error) *AppError {
	return &AppError{Code: 500, Message: message, Err: fmt.Errorf("%w: %v", ErrInternalServer, err)}
}
