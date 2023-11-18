package httperror

import (
	"net/http"
	"testing"
)

func TestSetOutputFormat(t *testing.T) {
	var isOk bool

	rollback := outputFormat
	defer func() {
		outputFormat = rollback
	}()

	SetOutputFormat(func(w http.ResponseWriter, message string, status int, code string) http.ResponseWriter {
		isOk = true
		return w
	})

	outputFormat(nil, "", 0, "")

	if !isOk {
		t.Error("SetOutputFormat() should set outputFormat")
	}
}

func TestGetInternalServerErrorMessage(t *testing.T) {
	if GetInternalServerErrorMessage() != internalServerErrorMessage {
		t.Error("GetInternalServerErrorMessage() should return internalServerErrorMessage")
	}
}

func TestSetInternalServerErrorMessage(t *testing.T) {
	rollback := internalServerErrorMessage
	defer func() {
		internalServerErrorMessage = rollback
	}()

	SetInternalServerErrorMessage("message")

	if GetInternalServerErrorMessage() != "message" {
		t.Error("SetInternalServerErrorMessage() should set internalServerErrorMessage")
	}
}
