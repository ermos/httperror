package httperror

import (
	"errors"
	"net/http"
)

// Middleware is a middleware that recovers from panics and translates the panic to HTTP errors.
func Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				message := internalServerErrorMessage
				code := "INTERNAL_SERVER_ERROR"
				status := http.StatusInternalServerError

				if rvr == http.ErrAbortHandler {
					// we don't recover http.ErrAbortHandler so the response
					// to the client is aborted, this should not be logged
					panic(rvr)
				}

				if err, isErr := rvr.(error); isErr {
					var localErr Error

					if errors.As(err, &localErr) {
						message = localErr.Error()
						status = localErr.Status()
						code = localErr.Code()
					} else {
						message = err.Error()
					}
				} else if text, isText := rvr.(string); isText {
					message = text
				}

				w = outputFormat(w, message, status, code)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
