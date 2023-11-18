package httperror

import (
	"context"
	"testing"
)

func TestErrorStructMessage(t *testing.T) {
	err := Error{
		message: "message",
	}

	if err.Error() != "message" {
		t.Error("Error() should return the message")
	}
}

func TestErrorStructCode(t *testing.T) {
	err := Error{
		code: "code",
	}

	if err.Code() != "code" {
		t.Error("Code() should return the code")
	}
}

func TestErrorStructStatus(t *testing.T) {
	err := Error{
		status: 404,
	}

	if err.Status() != 404 {
		t.Error("Status() should return the status")
	}
}

func TestErrorStructContext(t *testing.T) {
	ctx := context.Background()

	err := Error{
		ctx: ctx,
	}

	if err.Context() != ctx {
		t.Error("Context() should return the context")
	}
}
