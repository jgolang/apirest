package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetRouteVarValueString ...
func (api APIRest) GetRouteVarValueString(urlVarName string, r *http.Request) (string, error) {
	vars := api.Vars(r)
	value := vars[urlVarName]
	if value == "" {
		return value, fmt.Errorf("The route var %v has not been obtained", urlVarName)
	}
	return value, nil
}

// GetRouteVarValueInt ...
func (api APIRest) GetRouteVarValueInt(urlVarName string, r *http.Request) (int, error) {
	vars := api.Vars(r)
	value, err := strconv.Atoi(vars[urlVarName])
	if err != nil {
		return value, err
	}
	return value, nil
}

// GetRouteVarValueInt64 ...
func (api APIRest) GetRouteVarValueInt64(urlVarName string, r *http.Request) (int64, error) {
	vars := api.Vars(r)
	value, err := strconv.ParseInt(vars[urlVarName], 10, 64)
	if err != nil {
		return value, err
	}
	return value, nil
}

// GetRouteVarValueFloat64 ...
func (api APIRest) GetRouteVarValueFloat64(urlVarName string, r *http.Request) (float64, error) {
	vars := api.Vars(r)
	value, err := strconv.ParseFloat(vars[urlVarName], 64)
	if err != nil {
		return value, err
	}
	return value, nil
}

// GetRouteVarValueBool ...
func (api APIRest) GetRouteVarValueBool(urlVarName string, r *http.Request) (bool, error) {
	vars := api.Vars(r)
	value, err := strconv.ParseBool(vars[urlVarName])
	if err != nil {
		return false, err
	}
	return value, nil
}

type contextKey int

const (
	varsKey contextKey = iota
)

// Vars returns the route variables for the current request, if any.
func (api APIRest) Vars(r *http.Request) map[string]string {
	if rv := r.Context().Value(varsKey); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

// GetHeaderValueString doc ...
func (api APIRest) GetHeaderValueString(key string, r *http.Request) (string, error) {
	value := r.Header.Get(key)
	if value == "" {
		return value, fmt.Errorf("The %v key header has not been obtained", key)
	}
	return value, nil
}

// GetHeaderValueInt doc ...
func (api APIRest) GetHeaderValueInt(key string, r *http.Request) (int, error) {
	value, err := strconv.Atoi(r.Header.Get(key))
	if err != nil {
		return value, err
	}
	return value, nil
}

// GetHeaderValueInt64 doc ...
func (api APIRest) GetHeaderValueInt64(key string, r *http.Request) (int64, error) {
	value, err := strconv.ParseInt(r.Header.Get(key), 10, 64)
	if err != nil {
		return value, err
	}
	return value, nil
}

// GetHeaderValueFloat64 doc ...
func (api APIRest) GetHeaderValueFloat64(key string, r *http.Request) (float64, error) {
	value, err := strconv.ParseFloat(r.Header.Get(key), 64)
	if err != nil {
		return value, err
	}
	return value, nil
}

// GetHeaderValueBool doc ...
func (api APIRest) GetHeaderValueBool(key string, r *http.Request) (bool, error) {
	value, err := strconv.ParseBool(r.Header.Get(key))
	if err != nil {
		return value, err
	}
	return value, nil
}

// GetQueryParamValueString ...
func (api APIRest) GetQueryParamValueString(queryParamName string, r *http.Request) (string, error) {
	value := r.URL.Query().Get(queryParamName)
	if value == "" {
		return value, fmt.Errorf("The query parameter %v has not been obtained", queryParamName)
	}
	return value, nil
}

// GetQueryParamValueInt ...
func (api APIRest) GetQueryParamValueInt(queryParamName string, r *http.Request) (int, error) {
	value, err := strconv.Atoi(r.URL.Query().Get(queryParamName))
	if err != nil {
		return value, err
	}
	return value, nil
}

// GetQueryParamValueInt64 ...
func (api APIRest) GetQueryParamValueInt64(queryParamName string, r *http.Request) (int64, error) {
	value, err := strconv.ParseInt(r.URL.Query().Get(queryParamName), 10, 64)
	if err != nil {
		return value, err
	}
	return value, nil
}

// GetQueryParamValueFloat64 ...
func (api APIRest) GetQueryParamValueFloat64(queryParamName string, r *http.Request) (float64, error) {
	value, err := strconv.ParseFloat(r.URL.Query().Get(queryParamName), 64)
	if err != nil {
		return value, err
	}
	return value, nil
}

// GetQueryParamValueBool ...
func (api APIRest) GetQueryParamValueBool(queryParamName string, r *http.Request) (bool, error) {
	value, err := strconv.ParseBool(r.URL.Query().Get(queryParamName))
	if err != nil {
		return false, err
	}
	return value, nil
}

// UnmarshalBody doc ...
func (api APIRest) UnmarshalBody(v interface{}, r *http.Request) error {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyRequest, v)
	if err != nil {
		return err
	}
	return nil
}
