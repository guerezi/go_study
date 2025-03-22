package errors

import "fmt"

// Implementa unwarp and error
type Error struct {
	Message string
	Code    int
	Next    error
}

const (
	ErrorCodeNotFound int = iota
	ErrorCodeInvalid
)

// Error implements error.
func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

func (e *Error) Unwarp() error {
	return e.Next
}

func NewError(message string, code int, next error) error {
	return &Error{
		Message: message,
		Code:    code,
		Next:    next,
	}
}

var _ error = new(Error)

// var _ Unwrap = new(Error)
