package apirest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ToolBasic doc ...
type ToolBasic struct {
	R *http.Request
}

//GetURLParamString ...
func (t ToolBasic) GetURLParamString(paramName string) (string, Response) {
	params := mux.Vars(t.R)
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
func (t ToolBasic) GetURLParamInt(paramName string) (int, Response) {
	params := mux.Vars(t.R)
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
func (t ToolBasic) GetURLParamInt64(paramName string) (int64, Response) {

	params := mux.Vars(t.R)
	param, err := strconv.ParseInt(params[paramName], 10, 64)
	if Checkp(err) {
		return 0, Error{
			Title:   "Se esperaba un parametro tipo int 64",
			Message: fmt.Sprintf("No se ha obtenido el parametro %v", paramName),
		}
	}
	return param, nil

}

// GetQueryParamInt64 ...
func (t ToolBasic) GetQueryParamInt64(paramName string) (int64, Response) {

	param, err := strconv.ParseInt(t.R.URL.Query().Get(paramName), 10, 64)
	if Checkp(err) {
		return 0, Error{
			Title:   "Se esperaba un parametro tipo int 64",
			Message: fmt.Sprintf("No se ha obtenido el parametro %v", paramName),
		}
	}
	return param, nil

}

// GetQueryParamString ...
func (t ToolBasic) GetQueryParamString(paramName string) (string, Response) {

	param := t.R.URL.Query().Get(paramName)
	return param, nil

}
