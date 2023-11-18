package httperror

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	tests := []struct {
		name         string
		panicValue   interface{} // Value to panic with
		expectedCode int         // Expected HTTP status code
		expectedBody string      // Expected response body
	}{
		{
			name:         "OK",
			panicValue:   nil,
			expectedCode: http.StatusOK,
			expectedBody: "",
		},
		{
			name:         "Internal Error",
			panicValue:   errors.New("test error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"message":"an error occurred","code":"INTERNAL_SERVER_ERROR"}`,
		},
		{
			name:         "Error with httperror.Error",
			panicValue:   NewBadRequest("bad request"),
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message":"bad request"}`,
		},
		{
			name:         "Error with text",
			panicValue:   "oops",
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"message":"an error occurred","code":"INTERNAL_SERVER_ERROR"}`,
		},
		{
			name:         "Abort",
			panicValue:   http.ErrAbortHandler,
			expectedCode: http.StatusOK,
			expectedBody: "",
		},
	}

	for _, tt := range tests {
		dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if tt.panicValue != nil {
				panic(tt.panicValue)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			// Create a recorder to capture the response
			recorder := httptest.NewRecorder()

			// Create a request to pass to the middleware
			req := httptest.NewRequest("GET", "/", nil)

			// Set up the middleware
			middleware := Middleware(dummyHandler)

			// Execute the middleware with the panic value
			func() {
				defer func() {
					if rvr := recover(); rvr != nil {
						if rvr != http.ErrAbortHandler {
							t.Errorf("unexpected panic value: %v", rvr)
						}
					}
				}()

				req = httptest.NewRequest("GET", "/", nil) // Reset request after panic
				middleware.ServeHTTP(recorder, req)
			}()

			// Check the response status code
			if recorder.Code != tt.expectedCode {
				t.Errorf("expected status code %d, got %d", tt.expectedCode, recorder.Code)
			}

			// Check the response body
			if recorder.Body.String() != tt.expectedBody {
				t.Errorf("expected body %s, got %s", tt.expectedBody, recorder.Body.String())
			}
		})
	}
}
