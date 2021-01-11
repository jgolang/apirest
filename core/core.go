package core

import "net/http"

// New doc ...
func New(v APIRequestValidater, f APIResponseFormatter, r APIResponder) *APIRest {
	return &APIRest{
		requestValidator: v,
		formatter:        f,
		responder:        r,
	}
}

// APIRest doc ...
type APIRest struct {
	requestValidator APIRequestValidater
	formatter        APIResponseFormatter
	responder        APIResponder
}

// Respond ...
func (api *APIRest) Respond(data ResponseData, w http.ResponseWriter) {
	responseFormatted := api.formatter.Format(data)
	api.responder.Respond(responseFormatted, w)
}

// RegisterNewAPIResponseFormatter doc ...
func (api *APIRest) RegisterNewAPIResponseFormatter(f APIResponseFormatter) {
	api.formatter = f
}

// RegisterNewAPIResponder doc ...
func (api *APIRest) RegisterNewAPIResponder(r APIResponder) {
	api.responder = r
}

// RegisterNewAPIRequest doc ...
func (api *APIRest) RegisterNewAPIRequest(v APIRequestValidater) {
	api.requestValidator = v
}
