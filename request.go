package apigolang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Request ...
type Request struct {
	HTTPReq    *http.Request
	JSONStruct interface{}
}

//UnmarshalBody ...
func (request *Request) UnmarshalBody() Response {
	//  Read body JSON
	bodyRequest, err := ioutil.ReadAll(request.HTTPReq.Body)
	if err != nil {
		log.Println(err)
		return Error{
			Title:   "Estructura JSON invalida",
			Message: "Se requiere el cuerpo del objeto en el request",
		}
	}

	//  Unmarshal JSON to golang struct and validate
	unmErr := json.Unmarshal(bodyRequest, &request.JSONStruct)
	if unmErr != nil {
		log.Println(unmErr)
		return Error{
			Title:   "Estructura JSON invalida",
			Message: "No se ha leído la estrucutura...",
		}
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
	if err != nil {
		log.Println(err)
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
	if err != nil {
		log.Println(err)
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
	if err != nil {
		log.Println(err)
		return 0, Error{
			Title:   "Se esperaba un parametro tipo int 64",
			Message: fmt.Sprintf("No se ha obtenido el parametro %v", paramName),
		}
	}
	return param, nil
}
