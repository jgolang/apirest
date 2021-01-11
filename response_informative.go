package apirest

import (
	"net/http"

	"github.com/jgolang/apirest/core"
)

var (
	// DefaultInfoTitle doc ...
	DefaultInfoTitle = "Information!"
	// DefaultInfoMessage doc ..
	DefaultInfoMessage = "The request has been successful!"
	// InformativeType info response type the value is "info"
	InformativeType core.ResponseType = "info"
)

// Informative info response type the value is "info"
type Informative core.ResponseData

// Send ...
func (info Informative) Send(w http.ResponseWriter) {

	info.ResponseType = InformativeType

	if info.Title == "" {
		info.Title = DefaultInfoTitle
	}

	if info.Message == "" {
		info.Message = DefaultInfoMessage
	}

	if info.StatusCode == 0 {
		info.StatusCode = http.StatusOK
	}

	api.Respond(core.ResponseData(info), w)
}
