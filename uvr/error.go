package uvr

import (
	"errors"
	"fmt"
)

// NewErrorf returns an error from a format sring
func NewErrorf(format string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(format, a))
}

// NewError returns an error from a string
func NewError(message string) error {
	return errors.New(message)
}
