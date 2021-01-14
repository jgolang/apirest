package core

import "net/http"

// APIResponseFormatter doc ...
type APIResponseFormatter interface {
	Format(ResponseData) *ResponseFormatted
}

// APIResponder doc ...
type APIResponder interface {
	Respond(*ResponseFormatted, http.ResponseWriter)
}

// APISecurity doc ...
type APISecurity interface {
	ValidateBasicToken(token string) bool
	ValidateBearerToken(token string) bool
	ValidateCustomToken(func(string) bool) bool
}
