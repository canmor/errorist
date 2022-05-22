package errorist

import (
	"errors"
	"fmt"
)

type parameterError struct {
	error
	params []any
}

func newParameterError(err error, params ...any) *parameterError {
	return &parameterError{error: err, params: params}
}

func (e *parameterError) Error() string {
	return fmt.Sprintf("%v, params %v", e.error, e.params)
}

func (e *parameterError) Unwrap() error {
	return e.error
}

// UnwrapParams Unwrap parameterError from error and return its parameter slice
// err passed in should wrap with type `parameterError`, otherwise it returns nil slice
func UnwrapParams(err error) []any {
	var p *parameterError
	if errors.As(err, &p) {
		return p.params
	}
	return nil
}

func WrapParams(err error, params ...any) error {
	return newParameterError(err, params...)
}
