package apirest

import (
	"net/http"

	"github.com/jgolang/apirest/core"
)

var apiRest = core.New(
	RequestValidator{},
	ResponseFormatter{},
	Responder{},
	&mapMethods,
)

// RegisterNewAPIFormatter doc ...
func RegisterNewAPIFormatter(f core.APIResponseFormatter) {
	apiRest.RegisterNewAPIResponseFormatter(f)
}

// RegisterNewAPIResponder doc ...
func RegisterNewAPIResponder(f core.APIResponder) {
	apiRest.RegisterNewAPIResponder(f)
}

// RegisterNewAPIRequestValidator doc ...
func RegisterNewAPIRequestValidator(v core.APIRequestValidater) {
	apiRest.RegisterNewAPIRequestValidator(v)
}

// AddNewMapMethod doc ...
func AddNewMapMethod(key string, methods []string) {
	apiRest.AddMapMethod(key, methods)
}

var mapMethods core.MapMethods

func init() {
	mapMethods[availableRequestbodymiddleware] = []string{
		http.MethodPost,
		http.MethodGet,
		http.MethodPut,
		http.MethodPatch,
	}
}
