package apirest

import (
	"encoding/json"
	"net/http"
)

// ResponseType contains all the response types identiers
type ResponseType string

var (
	// ErrorType error response type the value is "error"
	ErrorType ResponseType = "error"
	// WarningType warning response type the value is "warning"
	WarningType ResponseType = "warning"
	// InformativeType info response type the value is "info"
	InformativeType ResponseType = "info"
	// SuccessType success response type the value is "success"
	SuccessType ResponseType = "success"
)

// Response doc ...
type Response interface {
	SendResponse(w http.ResponseWriter)
}

// ResponseBody response body structure
// contains the info section, with the response type and the messages for users
// and the content section, with the required data for the request
type ResponseBody struct {
	Info    ResponseInfo `json:"info"`
	Content interface{}  `json:"responseContent,omitempty"`
}

// ResponseInfo response body info section
type ResponseInfo struct {
	Type      ResponseType `json:"type"`
	Title     string       `json:"title,omitempty"`
	Message   string       `json:"message,omitempty"`
	Action    string       `json:"action,omitempty"`
	SessionID string       `json:"sessionId,omitempty"`
}

// ResponseData ...
type ResponseData struct {
	Type       ResponseType
	Title      string
	Message    string
	Action     string
	StatusCode int
	SessionID  string
	Content    interface{}
}

// SendResponse ...
func (response ResponseData) SendResponse(w http.ResponseWriter) {
	// set response headers and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	info := ResponseInfo{
		Type:      response.Type,
		Title:     response.Title,
		Message:   response.Message,
		Action:    response.Action,
		SessionID: response.SessionID,
	}

	responseBody := ResponseBody{
		Info:    info,
		Content: response.Content,
	}

	json.NewEncoder(w).Encode(responseBody)

}

// NewSuccessResponse ...
func NewSuccessResponse(title, message string, content interface{}) Response {
	return ResponseData{
		Title:      title,
		Message:    message,
		Content:    content,
		StatusCode: 200,
		Type:       SuccessType,
	}
}

// NewErrorResponse ...
func NewErrorResponse(title, message string, content interface{}) Response {
	return ResponseData{
		Title:      title,
		Message:    message,
		Content:    content,
		StatusCode: 400,
		Type:       ErrorType,
	}
}

// NewWarningResponse ...
func NewWarningResponse(title, message string, content interface{}) Response {
	return ResponseData{
		Title:      title,
		Message:    message,
		Content:    content,
		StatusCode: 200,
		Type:       WarningType,
	}
}

// NewInfoResponse ...
func NewInfoResponse(title, message string, content interface{}) Response {
	return ResponseData{
		Title:      title,
		Message:    message,
		Content:    content,
		StatusCode: 200,
		Type:       InformativeType,
	}
}

// SendSuccess ...
func SendSuccess(title, message string, w http.ResponseWriter) {
	response := NewSuccessResponse(title, message, nil)
	response.SendResponse(w)
}

// SendError ...
func SendError(title, message string, w http.ResponseWriter) {
	response := NewErrorResponse(title, message, nil)
	response.SendResponse(w)
}
