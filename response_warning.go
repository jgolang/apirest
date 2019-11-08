package apirest

import "net/http"

// Warning warning response type the value is "warning"
type Warning struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	Content    interface{}
}

// SetResponse warning ...
func (warning Warning) setResponse() (response ResponseData) {
	response = ResponseData{
		Title:      warning.Title,
		Message:    warning.Message,
		StatusCode: 200,
		Type:       WarningType,
		Action:     warning.Action,
		Content:    warning.Content,
	}
	if warning.StatusCode != 0 {
		response.StatusCode = warning.StatusCode
	}
	return
}

// SetWarningResponse ...
func SetWarningResponse(title, message string, w http.ResponseWriter) {
	warning := Warning{
		Title:   title,
		Message: message,
	}
	SendResponse(warning, w)
}
