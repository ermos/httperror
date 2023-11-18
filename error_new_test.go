package httperror

import (
	"context"
	"errors"
	"testing"
)

func TestNewError(t *testing.T) {
	err := NewError("message", 404)

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewError() should return an httperror.Error")
	}

	if goodType.Status() != 404 {
		t.Error("Status() should return the status")
	}
}

func TestNewErrorWithString(t *testing.T) {
	err := NewError("message", 404)

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewError() should return an httperror.Error")
	}

	if goodType.Error() != "message" {
		t.Error("Error() should return \"message\"")
	}
}

func TestNewErrorWithError(t *testing.T) {
	err := NewError(errors.New("message"), 404)

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewError() should return an httperror.Error")
	}

	if goodType.Error() != "message" {
		t.Error("Error() should return \"message\"")
	}
}

func TestNewErrorWithCustomType(t *testing.T) {
	err := NewError(struct {
		message string
	}{
		message: "message",
	}, 404)

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewError() should return an httperror.Error")
	}

	if goodType.Error() != "{message}" {
		t.Errorf("Error() should return \"{message}\", %v given", goodType.Error())
	}
}

func TestNewErrorWithCodeOption(t *testing.T) {
	err := NewError("message", 404, WithCode("code"))

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewError() should return an httperror.Error")
	}

	if goodType.Code() != "code" {
		t.Error("Code() should return \"code\"")
	}
}

func TestNewErrorWithContextOption(t *testing.T) {
	ctx := context.Background()

	err := NewError("message", 404, WithContext(ctx))

	var goodType Error

	if !errors.As(err, &goodType) {
		t.Error("NewError() should return an httperror.Error")
	}

	if goodType.Context() != ctx {
		t.Error("Context() should return the context")
	}
}
