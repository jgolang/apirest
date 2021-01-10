package core

import "net/http"

// New doc ...
func New(f APIFormatter, r APIResponder) *APIRest {
	return &APIRest{
		formatter: f,
		responder: r,
	}
}

// APIRest doc ...
type APIRest struct {
	formatter APIFormatter
	responder APIResponder
}

// Respond ...
func (api APIRest) Respond(data ResponseData, w http.ResponseWriter) {
	responseFormatted := api.formatter.Format(data)
	api.responder.Respond(responseFormatted, w)
}
