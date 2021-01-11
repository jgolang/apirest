package core

import "net/http"

// APIRequestValidater doc ...
type APIRequestValidater interface {
	Validate(*http.Request) (*RequestData, error)
}
