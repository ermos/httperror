package httperror

import (
	"context"
	"fmt"
)

// ErrorOption is a function that modifies an error.
type ErrorOption func(error *Error)

// NewError creates a new error with the given message and status code.
func NewError(err interface{}, status int, options ...ErrorOption) error {
	var result Error

	result.status = status

	if readErr, isErr := err.(error); isErr {
		result.message = readErr.Error()
	} else if readText, isText := err.(string); isText {
		result.message = readText
	} else {
		result.message = fmt.Sprintf("%v", err)
	}

	for _, opt := range options {
		opt(&result)
	}

	return result
}

// WithCode sets the code of the error.
func WithCode(code string) ErrorOption {
	return func(result *Error) {
		result.code = code
	}
}

// WithContext sets the context of the error.
func WithContext(ctx context.Context) ErrorOption {
	return func(result *Error) {
		result.ctx = ctx
	}
}
