package apirest

import (
	"encoding/base64"
	"net/http"
	"strings"
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

func validate(username, password string) bool {
	if username == "test" && password == "test" {
		return true
	}
	return false
}
