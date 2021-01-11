package apirest

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jgolang/apirest/core"
	"github.com/jgolang/log"
)

// MiddlewaresChain provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters
var MiddlewaresChain = core.MiddlewaresChain

// BasicAuth ...
func BasicAuth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			Error{Title: "Unauthorized!", StatusCode: 401}.Send(w)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			Error{Title: "Unauthorized!", StatusCode: 401}.Send(w)
			return
		}

		next(w, r)

	}
}

// RequestHeaderJSON validate header Content-Type, is required and equal to application/json
func RequestHeaderJSON(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		contentType := r.Header.Get("Content-Type")

		if len(contentType) == 0 {
			Error{Message: "No content-type!"}.Send(w)
			return
		}

		if contentType != "application/json" {
			Error{Message: "Content-Type not is JSON!"}.Send(w)
			return
		}

		next.ServeHTTP(w, r)

	}

}

// RequestHeaderSession doc ...
func RequestHeaderSession(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.Header.Get("SessionId")
		w.Header().Set("SessionId", sessionID)
		next.ServeHTTP(w, r)
	}
}

const availableRequestbodymiddleware = "availableRequestbodymiddleware"

// RequestBody doc ...
var RequestBody = NewRequestBodyMiddleware(availableRequestbodymiddleware)

// NewRequestBodyMiddleware doc ...
func NewRequestBodyMiddleware(keyListMethods string) core.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if api.ValidateMethods(keyListMethods, r.Method) {
				requestData, err := api.ValidateRequest(r)
				if err != nil {
					log.Error(err)
					Error{}.Send(w)
					return
				}
				b, valid := requestData.Data.(json.RawMessage)
				if !valid {
					log.Error("Invalid Content...")
					Error{}.Send(w)
					return
				}
				r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}
			next.ServeHTTP(w, r)
		}
	}
}
