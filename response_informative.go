package apirest

import "net/http"

var (
	// DefaultInfoTitle doc ...
	DefaultInfoTitle = "Information!"
	// DefaultInfoMessage doc ..
	DefaultInfoMessage = "The request has been successful!"
)

// Informative info response type the value is "info"
type Informative struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	SessionID  string
	Content    interface{}
}

// SetResponse informative ...
func (info Informative) setResponse() ResponseData {

	if info.Title == "" {
		info.Title = DefaultInfoTitle
	}

	if info.Message == "" {
		info.Message = DefaultInfoMessage
	}

	if info.StatusCode == 0 {
		info.StatusCode = http.StatusOK
	}

	return ResponseData{
		Title:      info.Title,
		Message:    info.Message,
		StatusCode: info.StatusCode,
		Type:       InformativeType,
		Action:     info.Action,
		SessionID:  info.SessionID,
		Content:    info.Content,
	}

}

// Send ...
func (info Informative) Send(w http.ResponseWriter) {
	SendResponse(info, w)
}
