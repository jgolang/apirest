package apirest

import (
	"encoding/json"
	"net/http"
)

// JSONRequest struct used to parse the request content section
type JSONRequest struct {
	Info    JSONRequestInfo `json:"info,omitempty"`
	Content json.RawMessage `json:"requestContent"`
}

// JSONRequestInfo request info section fields for encrypted requests
type JSONRequestInfo struct {
	DeviceUUID  string `json:"deviceUUID"`
	DeviceType  string `json:"deviceType"`
	OS          string `json:"os"`
	OSVersion   string `json:"osVersion"`
	OSTimezone  int    `json:"osTimezone"`
	AppLanguage string `json:"appLanguage"`
	AppVersion  string `json:"appVersion"`
	SessionID   string `json:"sessionId"`
}

//Request ...
type Request interface {
	GetSessionInfo(r *http.Request) (Request, Response)
	UnmarshalBody(r *http.Request, v interface{}) Response
}

// NewRequest ...
func NewRequest(req Request, r *http.Request) (Request, Response) {

	req, resp := req.GetSessionInfo(r)
	if resp.setResponse().StatusCode != 0 {
		return req, resp
	}

	return req, resp

}

// NewRequestBody ...
func NewRequestBody(req Request, r *http.Request, v interface{}) (Request, Response) {

	req, resp := req.GetSessionInfo(r)
	if resp.setResponse().StatusCode != 0 {
		return req, resp
	}

	resp = req.UnmarshalBody(r, v)
	if resp.setResponse().StatusCode != 0 {
		return req, resp
	}

	return req, resp

}
