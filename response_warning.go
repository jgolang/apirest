package apirest

import "net/http"

var (
	// DefaultWarningTitle doc ...
	DefaultWarningTitle = "Alert!"
	// DefaultWarningMessage doc ..
	DefaultWarningMessage = "The application has been successful but with potential problems!"
)

// Warning warning response type the value is "warning"
type Warning struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	SessionID  string
	Content    interface{}
}

// SetResponse warning ...
func (warning Warning) setResponse() ResponseData {

	if warning.Title == "" {
		warning.Title = DefaultWarningTitle
	}

	if warning.Message == "" {
		warning.Message = DefaultWarningMessage
	}

	if warning.StatusCode == 0 {
		warning.StatusCode = http.StatusOK
	}

	return ResponseData{
		Title:      warning.Title,
		Message:    warning.Message,
		StatusCode: warning.StatusCode,
		Type:       WarningType,
		Action:     warning.Action,
		SessionID:  warning.SessionID,
		Content:    warning.Content,
	}
}

// Send ...
func (warning Warning) Send(w http.ResponseWriter) {
	SendResponse(warning, w)
}
