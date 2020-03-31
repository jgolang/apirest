package apirest

import (
	"net/http"
)

// RequestBasic doc ...
type RequestBasic struct {
	JSONStruct interface{}
	SessionID  string
	UserID     string
	TraceID    string
	HTTPReq    *http.Request
}

// SessionIDPayload doc ...
var SessionIDPayload = "SessionID"

// UserIDPayload doc ...
var UserIDPayload = "UserID"

// TraceIDPayload doc ..
var TraceIDPayload = "event.TraceID"

//GetSessionID get session from user
func (request *RequestBasic) GetSessionID() Response {

	sessionID, response := GetHeaderValueString(SessionIDPayload, request.HTTPReq)
	if response != nil {
		resp := response.(Error)
		resp.Title = "Session info error!"
		resp.Message = "The session was not obtained"
		return response
	}

	request.SessionID = sessionID

	return nil
}

//GetUserID get id user session
func (request *RequestBasic) GetUserID() Response {

	userID, response := GetHeaderValueString(UserIDPayload, request.HTTPReq)
	if response != nil {
		resp := response.(Error)
		resp.Title = "Session info error!"
		resp.Message = "The user id was not obtained"
		return response
	}

	request.UserID = userID

	return nil
}

//GetTraceID doc
func (request *RequestBasic) GetTraceID() Response {

	traceID, response := GetHeaderValueString(TraceIDPayload, request.HTTPReq)
	if response != nil {
		resp := response.(Error)
		resp.Title = "Session info error!"
		resp.Message = "The trace id was not obtained"
		return response
	}

	request.TraceID = traceID

	return nil
}

// UnmarshalBody doc ...
func (request *RequestBasic) UnmarshalBody() Response {

	resp := UnmarshalBody(request.HTTPReq, request.JSONStruct)
	if resp != nil {
		return resp
	}

	return nil
}

//GetSessionInfo ..
func (request *RequestBasic) GetSessionInfo() Response {

	resp := request.GetSessionID()
	if resp != nil {
		return resp
	}

	resp = request.GetUserID()
	if resp != nil {
		return resp
	}

	resp = request.GetTraceID()
	if resp != nil {
		return resp
	}

	return nil

}

// GetRequestInfo ..
func (request *RequestBasic) GetRequestInfo() Response {

	resp := request.GetSessionInfo()
	if resp != nil {
		return resp
	}

	resp = request.UnmarshalBody()
	if resp != nil {
		return resp
	}

	return nil

}
