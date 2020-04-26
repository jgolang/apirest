package apirest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jgolang/log"
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

// GetHeaderValueString doc ...
func GetHeaderValueString(key string, r *http.Request) (string, Response) {
	value := r.Header.Get(key)
	if value == "" {
		return value, Error{
			Title:   "Error getting header!",
			Message: fmt.Sprintf("The %v key header has not been obtained", key),
		}
	}
	return value, nil
}

// GetHeaderValueInt doc ...
func GetHeaderValueInt(key string, r *http.Request) (int, Response) {
	value := r.Header.Get(key)
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return valueInt, Error{
			Title:   "Error getting header type Int!",
			Message: fmt.Sprintf("The %v key header has not been obtained", key),
		}
	}
	return valueInt, nil
}

// GetHeaderValueInt64 doc ...
func GetHeaderValueInt64(key string, r *http.Request) (int64, Response) {
	value := r.Header.Get(key)
	valueInt64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return valueInt64, Error{
			Title:   "Error getting header type Int64!",
			Message: fmt.Sprintf("The %v key header has not been obtained", key),
		}
	}
	return valueInt64, nil
}

// GetRouteVarValueString ...
func GetRouteVarValueString(urlVarName string, r *http.Request) (string, Response) {
	vars := mux.Vars(r)
	value := vars[urlVarName]
	if value == "" {
		return value, Error{
			Title:   "Error getting route var!",
			Message: fmt.Sprintf("The route var %v has not been obtained", urlVarName),
		}
	}
	return value, nil
}

// GetRouteVarValueInt ...
func GetRouteVarValueInt(urlVarName string, r *http.Request) (int, Response) {
	vars := mux.Vars(r)
	value, err := strconv.Atoi(vars[urlVarName]) //  Get user id
	if Check(err) {
		log.StackTrace(err)
		return 0, Error{
			Title:   "Error getting route var type Int",
			Message: fmt.Sprintf("The route var %v has not been obtained", urlVarName),
		}
	}
	return value, nil
}

// GetRouteVarValueInt64 ...
func GetRouteVarValueInt64(urlVarName string, r *http.Request) (int64, Response) {
	vars := mux.Vars(r)
	value, err := strconv.ParseInt(vars[urlVarName], 10, 64)
	if Check(err) {
		log.StackTrace(err)
		return value, Error{
			Title:   "Error getting route var type Int64",
			Message: fmt.Sprintf("The route var %v has not been obtained", urlVarName),
		}
	}
	return value, nil
}

// GetQueryParamValueString ...
func GetQueryParamValueString(queryParamName string, r *http.Request) (string, Response) {
	value := r.URL.Query().Get(queryParamName)
	if value == "" {
		return value, Error{
			Title:   "Error getting query param!",
			Message: fmt.Sprintf("The query parameter %v has not been obtained", queryParamName),
		}
	}

	return value, nil

}

// GetQueryParamValueInt ...
func GetQueryParamValueInt(queryParamName string, r *http.Request) (int, Response) {
	value, err := strconv.Atoi(r.URL.Query().Get(queryParamName))
	if Check(err) {
		log.StackTrace(err)
		return 0, Error{
			Title:   "Error getting query param type Int!",
			Message: fmt.Sprintf("The query parameter %v has not been obtained", queryParamName),
		}
	}
	return value, nil
}

// GetQueryParamValueInt64 ...
func GetQueryParamValueInt64(queryParamName string, r *http.Request) (int64, Response) {
	value, err := strconv.ParseInt(r.URL.Query().Get(queryParamName), 10, 64)
	if Check(err) {
		log.StackTrace(err)
		return 0, Error{
			Title:   "Error getting query param type Int64!",
			Message: fmt.Sprintf("The query parameter %v has not been obtained", queryParamName),
		}
	}
	return value, nil
}

// UnmarshalBody doc ...
func UnmarshalBody(v interface{}, r *http.Request) Response {

	//  Read body JSON
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if Check(err) {
		log.StackTrace(err)
		return Error{
			Title:   "Not read JSON struct!",
			Message: "Error when reading JSON structure",
		}
	}

	//  Unmarshal JSON to golang struct and validate
	err = json.Unmarshal(bodyRequest, &v)
	if Check(err) {
		log.StackTrace(err)
		return Error{
			Title:   "Invalid JSON struct!",
			Message: "Error when umarshal JSON structure",
		}
	}
	return nil

}
