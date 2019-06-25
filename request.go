package apigo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Request ...
type Request struct {
	HTTPReq    *http.Request
	JSONStruct interface{}
	SessionID  string
	UserID     int64
	TraceID    string
}

// NewRequest ...
func NewRequest(r *http.Request) (Request, Response) {
	request := Request{
		HTTPReq: r,
	}

	resp := request.GetSessionInfo()
	if resp.setResponse().StatusCode != 0 {
		return request, resp
	}

	return request, resp
}

// NewRequestBody ...
func NewRequestBody(r *http.Request, jsons interface{}) (Request, Response) {
	request := Request{
		HTTPReq:    r,
		JSONStruct: jsons,
	}

	resp := request.GetSessionInfo()
	if resp.setResponse().StatusCode != 0 {
		return request, resp
	}

	resp = request.UnmarshalBody()
	if resp.setResponse().StatusCode != 0 {
		return request, resp
	}

	return request, resp
}

//UnmarshalBody ...
func (request *Request) UnmarshalBody() Response {
	//  Read body JSON
	bodyRequest, err := ioutil.ReadAll(request.HTTPReq.Body)
	if Checkp(err) {
		return Error{
			Title:   "Estructura JSON invalida",
			Message: "Se requiere el cuerpo del objeto en el request",
		}
	}

	//  Unmarshal JSON to golang struct and validate
	unmErr := json.Unmarshal(bodyRequest, &request.JSONStruct)
	if Checkp(unmErr) {
		return Error{
			Title:   "Estructura JSON invalida",
			Message: "No se ha leído la estrucutura...",
		}
	}
	return nil
}

//GetSessionID get session from user
func (request *Request) GetSessionID() Response {
	request.SessionID = request.HTTPReq.Header.Get("SessionID")
	if request.SessionID == "" {
		return Error{
			Title:   "¡Error de session!",
			Message: "No se ha obtenido la session del usuario ",
		}
	}
	return nil
}

//GetUserID get id user session
func (request *Request) GetUserID() Response {
	userID, err := strconv.Atoi(request.HTTPReq.Header.Get("UserID"))
	if Checkp(err) {
		return Error{
			Title:   "¡Error de session!",
			Message: "No se ha obtenido el id del usuario",
		}
	}
	request.UserID = int64(userID)
	return nil
}

//GetTraceID doc
func (request *Request) GetTraceID() Response {
	request.TraceID = request.HTTPReq.Header.Get("event.TraceID")
	if request.TraceID == "" {
		return Error{
			Title:   "¡Error de session!",
			Message: "No se ha obtenido el id de la trasa",
		}
	}
	return nil
}

//GetSessionInfo ..
func (request *Request) GetSessionInfo() Response {
	resp := request.GetSessionID()
	if resp.setResponse().StatusCode != 0 {
		return resp
	}
	resp = request.GetUserID()
	if resp.setResponse().StatusCode != 0 {
		return resp
	}
	resp = request.GetTraceID()
	if resp.setResponse().StatusCode != 0 {
		return resp
	}
	return nil
}

//GetURLParamString ...
func (request *Request) GetURLParamString(paramName string) (string, Response) {
	params := mux.Vars(request.HTTPReq)
	param := params[paramName]
	if param == "" {
		return "", Error{
			Title:   "¡Error de tipo!",
			Message: fmt.Sprintf("No se ha obtenido el parametro %v", paramName),
		}
	}
	return param, nil
}

//GetURLParamInt ...
func (request *Request) GetURLParamInt(paramName string) (int, Response) {
	params := mux.Vars(request.HTTPReq)
	param, err := strconv.Atoi(params[paramName]) //  Get user id
	if Checkp(err) {
		return 0, Error{
			Title:   "Se esperaba un parametro tipo int",
			Message: fmt.Sprintf("No se ha obtenido el parametro %v", paramName),
		}
	}
	return param, nil
}

//GetURLParamInt64 ...
func (request *Request) GetURLParamInt64(paramName string) (int64, Response) {
	params := mux.Vars(request.HTTPReq)
	param, err := strconv.ParseInt(params[paramName], 10, 64) //  Get user id
	if Checkp(err) {
		return 0, Error{
			Title:   "Se esperaba un parametro tipo int 64",
			Message: fmt.Sprintf("No se ha obtenido el parametro %v", paramName),
		}
	}
	return param, nil
}

// GetQueryParamInt64 ...
func (request *Request) GetQueryParamInt64(paramName string) (int64, Response) {
	param, err := strconv.ParseInt(request.HTTPReq.URL.Query().Get(paramName), 10, 64) //  Get user id
	if Checkp(err) {
		return 0, Error{
			Title:   "Se esperaba un parametro tipo int 64",
			Message: fmt.Sprintf("No se ha obtenido el parametro %v", paramName),
		}
	}
	return param, nil
}
