package apigo

import (
	"net/http"
)

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
