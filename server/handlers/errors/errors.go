package errors

import "fmt"

type Error struct {
	Message string
	Status  int
}

// Error implements error.
func (e Error) Error() string {
	return fmt.Sprintf("Error %d: %s ", e.Status, e.Message)
}

var _ error = new(Error)
