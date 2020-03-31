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

<<<<<<< HEAD
// JSONResponse response body structure
// contains the info section, with the response type and the messages for users
// and the content section, with the required data for the request
type JSONResponse struct {
	Info    JSONResponseInfo `json:"info"`
	Content json.RawMessage  `json:"responseContent,omitempty"`
=======
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
>>>>>>> 63a2b798ceea35356b9bc7e319156243075624d4
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
<<<<<<< HEAD

=======
>>>>>>> 63a2b798ceea35356b9bc7e319156243075624d4
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

<<<<<<< HEAD
	var jsonContent []byte

	var err error
	if response.Content != nil {
		jsonContent, err = json.Marshal(response.Content)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	jsonResponse := JSONResponse{
=======
	responseBody := ResponseBody{
>>>>>>> 63a2b798ceea35356b9bc7e319156243075624d4
		Info:    info,
		Content: response.Content,
	}

<<<<<<< HEAD
	json.NewEncoder(w).Encode(jsonResponse)

}
=======
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
>>>>>>> 63a2b798ceea35356b9bc7e319156243075624d4
