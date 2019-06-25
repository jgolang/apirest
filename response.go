package apigo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
	Content    interface{}
}

// SendResponse ...
func (data *ResponseData) SendResponse(w http.ResponseWriter) {
	// set response headers and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.StatusCode)

	// build response body info section
	info := ResponseInfo{
		Type:      data.Type,
		Title:     data.Title,
		Message:   data.Message,
		Action:    data.Action,
		SessionID: w.Header().Get("SessionId"),
	}

	var jsonContent []byte
	var err error

	if data.Content != nil {
		jsonContent, err = json.Marshal(data.Content)
		if err != nil {
			log.Println(err)
			ErrorResponse("Lo sentimos", "No es posible responder en este momento, favor intentar mas tarde...", w)
		}
	}

	// build response body
	responseBody := ResponseBody{
		Info:    info,
		Content: jsonContent,
	}

	json.NewEncoder(w).Encode(responseBody)

	log.Println(w) //	Log

	encryptedResponseJSON, _ := json.Marshal(responseBody)
	body := ioutil.NopCloser(bytes.NewBuffer(encryptedResponseJSON))

	log.Println(body) //	Log
}

// Response ...
type Response interface {
	setResponse() ResponseData
}

// Success success response type the value is "success"
type Success struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	Content    interface{}
}

// Error error response type the value is "error"
type Error struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	Content    interface{}
}

// Warning warning response type the value is "warning"
type Warning struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	Content    interface{}
}

// Informative info response type the value is "info"
type Informative struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	Content    interface{}
}

// SetResponse success ...
func (succes Success) setResponse() (response ResponseData) {
	response = ResponseData{
		Title:      succes.Title,
		Message:    succes.Message,
		StatusCode: 200,
		Type:       SuccessType,
		Action:     succes.Action,
		Content:    succes.Content,
	}
	if succes.StatusCode != 0 {
		response.StatusCode = succes.StatusCode
	}
	return
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

// SetResponse informative ...
func (info Informative) setResponse() (response ResponseData) {
	response = ResponseData{
		Title:      info.Title,
		Message:    info.Message,
		StatusCode: 200,
		Type:       InformativeType,
		Action:     info.Action,
		Content:    info.Content,
	}
	if info.StatusCode != 0 {
		response.StatusCode = info.StatusCode
	}
	return
}

// NewResponse ...
func NewResponse(response Response) ResponseData {
	return response.setResponse()
}

// NewErrorResponse ...
func NewErrorResponse(title, message string) Error {
	return Error{
		Title:   title,
		Message: message,
	}
}

// NewSuccessResponse ...
func NewSuccessResponse(title, message string) Success {
	return Success{
		Title:   title,
		Message: message,
	}
}

// SendResponse ...
func SendResponse(response Response, w http.ResponseWriter) {
	res := response.setResponse()
	res.SendResponse(w)
}

// ErrorResponse ...
func ErrorResponse(title, message string, w http.ResponseWriter) {
	err := Error{
		Title:   title,
		Message: message,
	}
	SendResponse(err, w)
}

// SuccesResponse ...
func SuccesResponse(title, message string, w http.ResponseWriter) {
	success := Success{
		Title:   title,
		Message: message,
	}
	SendResponse(success, w)
}

// SuccesContentResponse ...
func SuccesContentResponse(title, message string, content interface{}, w http.ResponseWriter) {
	success := Success{
		Title:   title,
		Message: message,
		Content: content,
	}
	SendResponse(success, w)
}

// InformativeResponse ...
func InformativeResponse(title, message string, w http.ResponseWriter) {
	info := Informative{
		Title:   title,
		Message: message,
	}
	SendResponse(info, w)
}

// WarningResponse ...
func WarningResponse(title, message string, w http.ResponseWriter) {
	warning := Warning{
		Title:   title,
		Message: message,
	}
	SendResponse(warning, w)
}
