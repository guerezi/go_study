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
	ErrorDataBase
)

// Error implements error.
func (e *Error) Error() string {
	return fmt.Sprintf("Error on Usecase: %d - %s (%s)", e.Code, e.Message, e.Next)
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
