package core

import "net/http"

// APIFormatter doc ...
type APIFormatter interface {
	Format(ResponseData) ResponseFormatted
}

// APIResponder ...
type APIResponder interface {
	Respond(ResponseFormatted, http.ResponseWriter)
}
