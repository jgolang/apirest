package apigolang

import (
	"encoding/json"
	"net/http"
)

// Middleware provides a convenient mechanism for filtering HTTP requests
// entering the application. It returns a new handler which may perform various
// operations and should finish by calling the next HTTP handler.
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// EncryptedBody struct used to parse the encrypted body
type EncryptedBody struct {
	Data string `json:"data"`
}

// Info request info section fields for encrypted requests
type Info struct {
	DeviceUUID  string `json:"deviceUUID"`
	DeviceType  string `json:"deviceType"`
	OS          string `json:"os"`
	OSVersion   string `json:"osVersion"`
	OSTimezone  int    `json:"osTimezone"`
	AppLanguage string `json:"appLanguage"`
	AppVersion  string `json:"appVersion"`
	SessionID   string `json:"sessionId"`
}

// RequestInfo struct used to parse the request info section
type RequestInfo struct {
	Info Info `json:"info"`
}

// RequestContent struct used to parse the request content section
type RequestContent struct {
	RequestContent json.RawMessage `json:"requestContent"`
}

// MiddlewaresChain provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
func MiddlewaresChain(mw ...Middleware) Middleware {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}

// RequestHeaderJson validate header Content-Type, is required and equal to application/json
func RequestHeaderJson(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		contentType := r.Header.Get("Content-Type")

		if len(contentType) == 0 {
			ErrorResponse("Petición inválida", "", w)
			return
		}

		if contentType != "application/json" {
			ErrorResponse("Campos requeridos", "", w)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func RequestHeaderSession(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID := r.Header.Get("SessionId")
		w.Header().Set("SessionId", sessionID)
		next.ServeHTTP(w, r)
	}
}
