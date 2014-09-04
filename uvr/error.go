package uvr

import (
    "errors"
    "fmt"
)

func NewErrorf(format string, a ...interface{}) error {
    return  errors.New(fmt.Sprintf(format, a))
}