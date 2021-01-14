package core

import (
	"encoding/json"
	"net/http"
)

// New doc ...
func New(v APIRequestValidater, f APIResponseFormatter, r APIResponder, s APISecurity, mapMethods *MapMethods) *API {
	return &API{
		requestValidator: v,
		formatter:        f,
		responder:        r,
		security:         s,
		MapMethods:       mapMethods,
	}
}

// API doc ...
type API struct {
	requestValidator APIRequestValidater
	formatter        APIResponseFormatter
	responder        APIResponder
	security         APISecurity
	MapMethods       *MapMethods
}

// MapMethods doc ...
type MapMethods map[string][]string

// Respond doc ...
func (api *API) Respond(data ResponseData, w http.ResponseWriter) {
	responseFormatted := api.formatter.Format(data)
	api.responder.Respond(responseFormatted, w)
}

// ValidateRequest doc ...
func (api *API) ValidateRequest(r *http.Request) (*RequestData, error) {
	return api.requestValidator.ValidateRequest(r)
}

// RegisterNewAPIResponseFormatter doc ...
func (api *API) RegisterNewAPIResponseFormatter(f APIResponseFormatter) {
	api.formatter = f
}

// RegisterNewAPIResponder doc ...
func (api *API) RegisterNewAPIResponder(r APIResponder) {
	api.responder = r
}

// ValidateCustomToken doc ...
func (api *API) ValidateCustomToken(token string, customValidator CustomTokenValidator) (json.RawMessage, bool) {
	return api.security.ValidateCustomToken(token, customValidator)
}

// RegisterNewAPIRequestValidator doc ...
func (api *API) RegisterNewAPIRequestValidator(v APIRequestValidater) {
	api.requestValidator = v
}

// RegisterNewAPISecurity doc ...
func (api *API) RegisterNewAPISecurity(s APISecurity) {
	api.security = s
}

// AddMapMethod doc ...
func (api *API) AddMapMethod(key string, methods []string) {
	mapMethods := *api.MapMethods
	mapMethods[key] = methods
	api.MapMethods = &mapMethods
}
