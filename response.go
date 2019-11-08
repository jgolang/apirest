package apirest

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// ResponseType contains all the response types identiers
type ResponseType string

const (
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
	SendAsError(w http.ResponseWriter)
	SendAsSuccess(w http.ResponseWriter)
	SendAsWarning(w http.ResponseWriter)
	SendAsInfo(w http.ResponseWriter)
}

// ResponseBody response body structure
// contains the info section, with the response type and the messages for users
// and the content section, with the required data for the request
type ResponseBody struct {
	Info    ResponseInfo    `json:"info"`
	Content json.RawMessage `json:"responseContent,omitempty"`
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
func SendResponse(response ResponseData, w http.ResponseWriter) {
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

	var jsonContent []byte

	var err error
	if response.Content != nil {
		jsonContent, err = json.Marshal(response.Content)
		if err != nil {
			logger := zap.S()
			defer logger.Sync()
			logger.Error(err)
			ErrorResponse("Lo sentimos", "No es posible responder en este momento, favor intentar mas tarde...", w)
		}
	}

	responseBody := ResponseBody{
		Info:    info,
		Content: jsonContent,
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

// SendAsSuccess doc ...
func (response ResponseData) SendAsSuccess(w http.ResponseWriter) {
	response.Type = SuccessType
	response.StatusCode = 200
	SendResponse(response, w)
}

// SetSuccessResponse doc ...
func SetSuccessResponse(title, message string, content interface{}, w http.ResponseWriter) {

	response := NewSuccessResponse(title, message, content)
	SendResponse(w)

}

// // NewResponse ...
// func NewResponse(response Response) ResponseData {
// 	return response.setResponse()
// }

// // SendResponse ...
// func SendResponse(response Response, w http.ResponseWriter) {
// 	res := response.setResponse()
// 	res.SendResponse(w)
// }
