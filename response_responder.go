package apirest

import (
	"encoding/json"
	"net/http"

	"github.com/jgolang/apirest/core"
	"github.com/jgolang/log"
)

// Responder doc ...
type Responder struct{}

// Respond doc ...
func (r Responder) Respond(response *core.ResponseFormatted, w http.ResponseWriter) {

	for key := range response.Headers {
		w.Header().Set(key, response.Headers[key])
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	err := json.NewEncoder(w).Encode(response.Data)
	if err != nil {
		log.Fatal(err)
	}

	return
}
