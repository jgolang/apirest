package apirest

import "net/http"

var (
	// DefaultErrorTitle doc ...
	DefaultErrorTitle = "Error response!"
	// DefaultErrorMessage doc ..
	DefaultErrorMessage = "The service has not completed the operation!"
)

// Error error response type the value is "error"
type Error struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	Content    interface{}
}

// SetResponse error ...
func (err Error) setResponse() ResponseData {

	if err.Title == "" {
		err.Title = DefaultErrorTitle
	}

	if err.Message == "" {
		err.Message = DefaultErrorMessage
	}

	if err.StatusCode == 0 {
		err.StatusCode = http.StatusBadRequest

	}

	return ResponseData{
		Title:      err.Title,
		Message:    err.Message,
		StatusCode: err.StatusCode,
		Type:       ErrorType,
		Action:     err.Action,
		Content:    err.Content,
	}

}

// Send ...
func (err Error) Send(w http.ResponseWriter) {
	SendResponse(err, w)
}
