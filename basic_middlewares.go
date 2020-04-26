package apirest

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jgolang/log"
)

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

// RequestBody doc ...
func RequestBody(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch {
			// Decode the request body to JSON
			jsonDecoder := json.NewDecoder(r.Body)
			var request JSONRequest
			err := jsonDecoder.Decode(&request)

			if err != nil {
				log.Error(err)
				Error{Message: "Empty body is required to use this method!"}.Send(w)
				return
			}

			r.Header.Set("Request-Content", string(request.Content))
			r.Body = ioutil.NopCloser(bytes.NewBuffer(request.Content))

		}

		next.ServeHTTP(w, r)

	}

}
