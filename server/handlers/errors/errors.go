package errors

import "fmt"

type ErrorHandler struct {
	Message string
	Status  int
}

// Error implements error.
func (e ErrorHandler) Error() string {
	return fmt.Sprintf("Error %d: %s ", e.Status, e.Message)
}

var _ error = new(ErrorHandler)
