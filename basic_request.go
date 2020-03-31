package apirest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// RequestBasic doc ...
type RequestBasic struct {
	JSONStruct interface{}
	SessionID  string
	UserID     string
	TraceID    string
	Tools      ToolBasic
}

// SessionIDPayload doc ...
var SessionIDPayload = "SessionID"

// UserIDPayload doc ...
var UserIDPayload = "UserID"

// TraceIDPayload doc ..
var TraceIDPayload = "event.TraceID"

//UnmarshalBody ...
func (request RequestBasic) UnmarshalBody(r *http.Request, v interface{}) Response {

	req := RequestBasic{
		JSONStruct: v,
	}

	req.Tools.R = r

	//  Read body JSON
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if Checkp(err) {
		return Error{
			Title:   "Estructura JSON invalida",
			Message: "Se requiere el cuerpo del objeto en el request",
		}
	}

	//  Unmarshal JSON to golang struct and validate
	unmErr := json.Unmarshal(bodyRequest, &req.JSONStruct)
	if Checkp(unmErr) {
		return Error{
			Title:   "Estructura JSON invalida",
			Message: "No se ha leído la estructura...",
		}
	}
	return nil
}

//GetSessionID get session from user
func (request *RequestBasic) GetSessionID(r *http.Request) Response {
	request.SessionID = r.Header.Get(SessionIDPayload)
	if request.SessionID == "" {
		return Error{
			Title:   "¡Error de session!",
			Message: "No se ha obtenido la session del usuario ",
		}
	}
	return nil
}

//GetUserID get id user session
func (request *RequestBasic) GetUserID(r *http.Request) Response {
	request.UserID = r.Header.Get(UserIDPayload)
	if request.UserID == "" {
		return Error{
			Title:   "¡Error de session!",
			Message: "No se ha obtenido el id del usuario",
		}
	}
	return nil
}

//GetTraceID doc
func (request *RequestBasic) GetTraceID(r *http.Request) Response {
	request.TraceID = r.Header.Get(TraceIDPayload)
	if request.TraceID == "" {
		return Error{
			Title:   "¡Error de session!",
			Message: "No se ha obtenido el id del evento",
		}
	}
	return nil
}

//GetSessionInfo ..
func (request RequestBasic) GetSessionInfo(r *http.Request) (Request, Response) {

	req := RequestBasic{}

	req.Tools.R = r

	resp := request.GetSessionID(r)
	if resp.setResponse().StatusCode != 0 {
		return nil, resp
	}
	resp = request.GetUserID(r)
	if resp.setResponse().StatusCode != 0 {
		return nil, resp
	}
	resp = request.GetTraceID(r)
	if resp.setResponse().StatusCode != 0 {
		return nil, resp
	}
	return req, nil

}
