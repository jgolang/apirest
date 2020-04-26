package apirest

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"go.mnc.gt/log"
)

// Middleware provides a convenient mechanism for filtering HTTP requests
// entering the application. It returns a new handler which may perform various
// operations and should finish by calling the next HTTP handler.
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// MiddlewaresChain provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
func MiddlewaresChain(mw ...Middleware) Middleware {

	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}

			res := httptest.NewRecorder()

			printAPIRequest(r)

			last(res, r)

			printAPIResponse(res)

		}

	}

}

func printAPIRequest(r *http.Request) {

	if prod {
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	bodyStr := buf.String()
	log.Infof("Request: %v %v\n", r.Method, r.RequestURI)
	log.Infof("Headers: %v\n", r.Header)
	log.Infof("Form: %v\n", r.Form.Encode())
	log.Infof("Body: %v\n", bodyStr)

}

func printAPIResponse(res *httptest.ResponseRecorder) {

	if prod {
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	bodyStr := buf.String()
	log.Infof("Response: %v\n", res.Result())
	log.Infof("Status Code: %v %v\n", res.Code, http.StatusText(res.Code))
	log.Infof("Headers: %v\n", res.Header())
	log.Infof("Body: %v\n", bodyStr)

}
