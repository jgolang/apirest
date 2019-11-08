package apirest

import "net/http"

// Error error response type the value is "error"
type Error struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	Content    interface{}
}

// SetResponse error ...
func (err Error) setResponse() (response ResponseData) {
	response = ResponseData{
		Title:      err.Title,
		Message:    err.Message,
		StatusCode: 400,
		Type:       ErrorType,
		Action:     err.Action,
		Content:    err.Content,
	}
	if err.StatusCode != 0 {
		response.StatusCode = err.StatusCode
	}
	return response
}

// NewErrorResponse ...
func NewErrorResponse(title, message string) Error {
	return Error{
		Title:   title,
		Message: message,
	}
}

// ErrorResponse ...
func ErrorResponse(title, message string, w http.ResponseWriter) {
	err := Error{
		Title:   title,
		Message: message,
	}
	SendResponse(err, w)
}
