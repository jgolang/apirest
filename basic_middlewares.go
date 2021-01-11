package apirest

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

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

// RequestBody doc ...
var RequestBody = NewRequestBodyMiddleware(PPPGMethodsKey)

// NewRequestBodyMiddleware doc ...
func NewRequestBodyMiddleware(keyListMethods string) core.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if api.ValidateMethods(keyListMethods, r.Method) {
				requestData, err := api.ValidateRequest(r)
				if err != nil {
					log.Error(err)
					Error{
						Title:   "Invalid request content",
						Message: "Request content empty json structure",
					}.Send(w)
					return
				}
				r.Body = ioutil.NopCloser(bytes.NewBuffer(requestData.RawBody))
			}
			next.ServeHTTP(w, r)
		}
	}
}

// ContentExtractor doc
func ContentExtractor(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestData, err := api.ValidateRequest(r)
		if err != nil {
			log.Error(err)
			Error{
				Title:   "Invalid request content",
				Message: "Request content empty json structure",
			}.Send(w)
			return
		}

		proxiedIPAddress := r.Header.Get("X-Forwarded-For")
		if proxiedIPAddress != "" {
			ips := strings.Split(proxiedIPAddress, ", ")
			proxiedIPAddress = ips[0]
		} else {
			proxiedIPAddress = r.RemoteAddr
		}

		prefixEventID := requestData.DeviceUUID
		if prefixEventID == "" {
			prefixEventID = proxiedIPAddress
		}

		eventID := fmt.Sprintf("%v:%v:%v", prefixEventID, time.Now().UnixNano(), r.RequestURI)
		LogRequest(r.Method, r.RequestURI, eventID, r.Form.Encode(), r.Header, requestData.RawBody)

		r.Header.Set("EventID", eventID)
		r.Header.Set("DeviceUUID", requestData.DeviceUUID)
		r.Header.Set("DeviceType", requestData.DeviceType)
		r.Header.Set("DeviceOS", requestData.DeviceOS)
		r.Header.Set("OSVersion", requestData.OSVersion)
		r.Header.Set("OSTimezone", requestData.OSTimezone)
		r.Header.Set("AppLanguage", requestData.AppLanguage)
		r.Header.Set("AppVersion", requestData.AppVersion)
		r.Header.Set("AppName", requestData.AppName)
		r.Header.Set("SessionID", requestData.SessionID)

		b, valid := requestData.Data.(json.RawMessage)
		if !valid {
			log.Error("Invalid content...")
			Error{
				Title:   "Invalid content",
				Message: "Invalid content json structure",
			}.Send(w)
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(b)) // Set content data
		rec := httptest.NewRecorder()

		next.ServeHTTP(rec, r)

		for k, v := range rec.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(rec.Code)
		w.Write(rec.Body.Bytes())
		go LogResponse(rec)
	}
}
