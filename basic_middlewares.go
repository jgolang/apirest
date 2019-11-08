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
			ErrorResponse("No autorizado", "", w)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			ErrorResponse("No autorizado", "", w)
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
			ErrorResponse("Petición inválida", "", w)
			return
		}

		// if contentType != "application/json" {
		// 	ErrorResponse("Campos requeridos", "", w)
		// 	return
		// }

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

// GetRequestBodyMiddleware doc ...
func GetRequestBodyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch {
			// Decode the request body to JSON
			jsonDecoder := json.NewDecoder(r.Body)
			var requestContent RequestContent

			parsingRequestContentError := jsonDecoder.Decode(&requestContent)

			if parsingRequestContentError != nil {
				ErrorResponse("Error interno", "", w)
				return
			}

			r.Header.Set("Request-Content", string(requestContent.RequestContent))
			r.Body = ioutil.NopCloser(bytes.NewBuffer(requestContent.RequestContent))
			log.Println(r)
		}
		next.ServeHTTP(w, r)
	}
}
