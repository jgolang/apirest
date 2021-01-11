package apirest

import (
	"net/http"

	"github.com/jgolang/apirest/core"
)

var api = core.New(
	RequestValidator{},
	ResponseFormatter{},
	Responder{},
	&mapMethods,
)

// RegisterNewAPIFormatter doc ...
func RegisterNewAPIFormatter(f core.APIResponseFormatter) {
	api.RegisterNewAPIResponseFormatter(f)
}

// RegisterNewAPIResponder doc ...
func RegisterNewAPIResponder(f core.APIResponder) {
	api.RegisterNewAPIResponder(f)
}

// RegisterNewAPIRequestValidator doc ...
func RegisterNewAPIRequestValidator(v core.APIRequestValidater) {
	api.RegisterNewAPIRequestValidator(v)
}

// AddNewMapMethod doc ...
func AddNewMapMethod(key string, methods []string) {
	api.AddMapMethod(key, methods)
}

var mapMethods core.MapMethods

func init() {
	mapMethods[AvailableRequestbodymiddleware] = []string{
		http.MethodPost,
		http.MethodGet,
		http.MethodPut,
		http.MethodPatch,
	}
}
