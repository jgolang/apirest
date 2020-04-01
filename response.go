package apirest

import (
	"encoding/json"
	"log"
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

// JSONResponse response body structure
// contains the info section, with the response type and the messages for users
// and the content section, with the required data for the request
type JSONResponse struct {
	Info    JSONResponseInfo `json:"info"`
	Content interface{}      `json:"responseContent,omitempty"`
}

// JSONResponseInfo response body info section
type JSONResponseInfo struct {
	Type      ResponseType `json:"type"`
	Title     string       `json:"title,omitempty"`
	Message   string       `json:"message,omitempty"`
	Action    string       `json:"action,omitempty"`
	SessionID string       `json:"sessionId,omitempty"`
}

// Response ...
type Response interface {
	setResponse() ResponseData
	Send(w http.ResponseWriter)
}

// SendResponse ...
func SendResponse(response Response, w http.ResponseWriter) {
	res := response.setResponse()
	res.SendResponse(w)
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

	info := JSONResponseInfo{
		Type:      response.Type,
		Title:     response.Title,
		Message:   response.Message,
		Action:    response.Action,
		SessionID: response.SessionID,
	}

	jsonResponse := JSONResponse{
		Info:    info,
		Content: response.Content,
	}

	err := json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		log.Fatal(err)
	}

	return

}
