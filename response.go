package apigolang

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
	Type      string `json:"type"`
	Title     string `json:"title,omitempty"`
	Message   string `json:"message,omitempty"`
	Action    string `json:"action,omitempty"`
	SessionID string `json:"sessionId,omitempty"`
}

// ResponseData ...
type ResponseData struct {
	Type       string
	Title      string
	Message    string
	Action     string
	StatusCode int
	Content    json.RawMessage
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

	// build response body
	responseBody := ResponseBody{
		Info:    info,
		Content: data.Content,
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
	Content    json.RawMessage
}

// Error error response type the value is "error"
type Error struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	Content    json.RawMessage
}

// Warning warning response type the value is "warning"
type Warning struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	Content    json.RawMessage
}

// Informative info response type the value is "info"
type Informative struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	Content    json.RawMessage
}

// SetResponse success ...
func (succes Success) setResponse() (response ResponseData) {
	response = ResponseData{
		Title:      succes.Title,
		Message:    succes.Message,
		StatusCode: 200,
		Type:       "Success",
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
		Type:       "Error",
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
		Type:       "Warning",
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
		Type:       "Informative",
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
