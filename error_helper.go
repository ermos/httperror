package httperror

import (
	"net/http"
)

// NewBadRequest returns a new error with status code 400.
func NewBadRequest(err interface{}, options ...ErrorOption) error {
	return NewError(err, http.StatusBadRequest, options...)
}

// NewInternalServerError returns a new error with status code 500.
func NewInternalServerError(err interface{}, options ...ErrorOption) error {
	return NewError(err, http.StatusInternalServerError, options...)
}

// NewNotFound returns a new error with status code 404.
func NewNotFound(err interface{}, options ...ErrorOption) error {
	return NewError(err, http.StatusNotFound, options...)
}

// NewUnauthorized returns a new error with status code 401.
func NewUnauthorized(err interface{}, options ...ErrorOption) error {
	return NewError(err, http.StatusUnauthorized, options...)
}

// NewForbidden returns a new error with status code 403.
func NewForbidden(err interface{}, options ...ErrorOption) error {
	return NewError(err, http.StatusForbidden, options...)
}

// NewConflict returns a new error with status code 409.
func NewConflict(err interface{}, options ...ErrorOption) error {
	return NewError(err, http.StatusConflict, options...)
}

// NewNotImplemented returns a new error with status code 501.
func NewNotImplemented(err interface{}, options ...ErrorOption) error {
	return NewError(err, http.StatusNotImplemented, options...)
}

// NewMethodNotAllowed returns a new error with status code 405.
func NewMethodNotAllowed(err interface{}, options ...ErrorOption) error {
	return NewError(err, http.StatusMethodNotAllowed, options...)
}

// NewUnprocessableEntity returns a new error with status code 422.
func NewUnprocessableEntity(err interface{}, options ...ErrorOption) error {
	return NewError(err, http.StatusUnprocessableEntity, options...)
}
