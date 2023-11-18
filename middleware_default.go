package httperror

import (
	"encoding/json"
	"log"
	"net/http"
)

type defaultErrorFormat struct {
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

func defaultOutputFormat(w http.ResponseWriter, message string, status int, code string) http.ResponseWriter {
	if status >= 500 {
		log.Printf("error happened. Err: %s", message)
		message = GetInternalServerErrorMessage()
		code = "INTERNAL_SERVER_ERROR"
	}

	response, err := json.Marshal(defaultErrorFormat{Message: message, Code: code})
	if err != nil {
		log.Printf("error happened in JSON marshal. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return w
	}

	w.WriteHeader(status)

	_, err = w.Write(response)
	if err != nil {
		log.Printf("error happened in writing response. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	return w
}
