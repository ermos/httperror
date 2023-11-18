package httperror

import (
	"errors"
	"testing"
)

func TestNewBadRequest(t *testing.T) {
	err := NewBadRequest("message")

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewBadRequest() should return an httperror.Error")
	}

	if goodType.Status() != 400 {
		t.Error("Status() should return 400")
	}
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("message")

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewInternalServerError() should return an httperror.Error")
	}

	if goodType.Status() != 500 {
		t.Error("Status() should return 500")
	}
}

func TestNewNotFound(t *testing.T) {
	err := NewNotFound("message")

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewNotFound() should return an httperror.Error")
	}

	if goodType.Status() != 404 {
		t.Error("Status() should return 404")
	}
}

func TestNewUnauthorized(t *testing.T) {
	err := NewUnauthorized("message")

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewUnauthorized() should return an httperror.Error")
	}

	if goodType.Status() != 401 {
		t.Error("Status() should return 401")
	}
}

func TestNewForbidden(t *testing.T) {
	err := NewForbidden("message")

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewForbidden() should return an httperror.Error")
	}

	if goodType.Status() != 403 {
		t.Error("Status() should return 403")
	}
}

func TestNewConflict(t *testing.T) {
	err := NewConflict("message")

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewConflict() should return an httperror.Error")
	}

	if goodType.Status() != 409 {
		t.Error("Status() should return 409")
	}
}

func TestNewNotImplemented(t *testing.T) {
	err := NewNotImplemented("message")

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewNotImplemented() should return an httperror.Error")
	}

	if goodType.Status() != 501 {
		t.Error("Status() should return 501")
	}
}

func TestNewMethodNotAllowed(t *testing.T) {
	err := NewMethodNotAllowed("message")

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewMethodNotAllowed() should return an httperror.Error")
	}

	if goodType.Status() != 405 {
		t.Error("Status() should return 405")
	}
}

func TestNewUnprocessableEntity(t *testing.T) {
	err := NewUnprocessableEntity("message")

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewUnprocessableEntity() should return an httperror.Error")
	}

	if goodType.Status() != 422 {
		t.Error("Status() should return 422")
	}
}
