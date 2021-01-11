package core

import "net/http"

// APIResponseFormatter doc ...
type APIResponseFormatter interface {
	Format(ResponseData) *ResponseFormatted
}

// APIResponder ...
type APIResponder interface {
	Respond(*ResponseFormatted, http.ResponseWriter)
}
