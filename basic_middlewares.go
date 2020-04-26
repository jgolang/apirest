package apirest

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"
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

			var reqStruct JSONRequest

			response := UnmarshalBody(reqStruct, r)
			if response != nil {
				response.Send(w)
				return
			}

			r.Header.Set("Request-Content", string(reqStruct.Content))
			r.Body = ioutil.NopCloser(bytes.NewBuffer(reqStruct.Content))

		}

		next.ServeHTTP(w, r)

	}

}
