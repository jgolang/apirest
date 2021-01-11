package apirest

import (
	"net/http"

	"github.com/jgolang/apirest/core"
)

// RequestValidator implementation
type RequestValidator struct{}

// ValidateRequest doc
func (v RequestValidator) ValidateRequest(r *http.Request) (*core.RequestData, error) {
	return nil, nil
}
