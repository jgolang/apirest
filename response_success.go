package apirest

import "net/http"

var (
	// DefaultSuccessTitle doc ...
	DefaultSuccessTitle = "Congratulations!"
	// DefaultSuccessMessage doc ..
	DefaultSuccessMessage = "The request has succeeded!"
)

// Success success response type the value is "success"
type Success struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	Content    interface{}
	ResponseData
}

// SetResponse success ...
func (success Success) setResponse() ResponseData {

	if success.Title == "" {
		success.Title = DefaultSuccessTitle
	}

	if success.Message == "" {
		success.Message = DefaultSuccessMessage
	}

	if success.StatusCode == 0 {
		success.StatusCode = http.StatusOK
	}

	return ResponseData{
		Title:      success.Title,
		Message:    success.Message,
		StatusCode: success.StatusCode,
		Type:       SuccessType,
		Action:     success.Action,
		Content:    success.Content,
	}

}

// Send ...
func (success Success) Send(w http.ResponseWriter) {
	SendResponse(success, w)
}
