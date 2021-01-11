package core

import "net/http"

// New doc ...
func New(v APIRequestValidater, f APIResponseFormatter, r APIResponder) *API {
	return &API{
		requestValidator: v,
		formatter:        f,
		responder:        r,
	}
}

// API doc ...
type API struct {
	requestValidator APIRequestValidater
	formatter        APIResponseFormatter
	responder        APIResponder
}

// Respond doc ...
func (api *API) Respond(data ResponseData, w http.ResponseWriter) {
	responseFormatted := api.formatter.Format(data)
	api.responder.Respond(responseFormatted, w)
}

// Validate doc ...
func (api *API) Validate(r *http.Request) (*RequestData, error) {
	return api.requestValidator.Validate(r)
}

// RegisterNewAPIResponseFormatter doc ...
func (api *API) RegisterNewAPIResponseFormatter(f APIResponseFormatter) {
	api.formatter = f
}

// RegisterNewAPIResponder doc ...
func (api *API) RegisterNewAPIResponder(r APIResponder) {
	api.responder = r
}

// RegisterNewAPIRequest doc ...
func (api *API) RegisterNewAPIRequest(v APIRequestValidater) {
	api.requestValidator = v
}
