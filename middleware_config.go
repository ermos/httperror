package httperror

import "net/http"

var outputFormat = defaultOutputFormat
var internalServerErrorMessage = "an error occurred"

func SetOutputFormat(method func(w http.ResponseWriter, message string, status int, code string) http.ResponseWriter) {
	outputFormat = method
}

func GetInternalServerErrorMessage() string {
	return internalServerErrorMessage
}

func SetInternalServerErrorMessage(message string) {
	internalServerErrorMessage = message
}
