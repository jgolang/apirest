package apirest

import (
	"net/http"

	"github.com/jgolang/apirest/core"
)

var (
	// DefaultErrorTitle doc ...
	DefaultErrorTitle = "Error response!"
	// DefaultErrorMessage doc ..
	DefaultErrorMessage = "The service has not completed the operation!"
	// ErrorType error response type the value is "error"
	ErrorType core.ResponseType = "error"
)

// Error error response type the value is "error"
type Error core.ResponseData

// Send ...
func (err Error) Send(w http.ResponseWriter) {

	err.ResponseType = ErrorType

	if err.Title == "" {
		err.Title = DefaultErrorTitle
	}

	if err.Message == "" {
		err.Message = DefaultErrorMessage
	}

	if err.StatusCode == 0 {
		err.StatusCode = http.StatusBadRequest
	}

	if err.ErrorCode == "0" {
		err.ErrorCode = "1"
	}

	api.Respond(core.ResponseData(err), w)
}
