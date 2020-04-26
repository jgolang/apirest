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

			for i, header := range res.Header() {
				w.Header()[i] = header
			}

			w.WriteHeader(res.Code)
			w.Write(res.Body.Bytes())

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
	log.Infof("Request: %v %v", r.Method, r.RequestURI)
	log.Infof("Headers: %v", r.Header)
	log.Infof("Form: %v", r.Form.Encode())
	log.Infof("Body: \n%v", bodyStr)

}

func printAPIResponse(res *httptest.ResponseRecorder) {

	if prod {
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	bodyStr := buf.String()
	log.Infof("Response: %v", res.Result())
	log.Infof("Status Code: %v %v", res.Code, http.StatusText(res.Code))
	log.Infof("Headers: %v", res.Header())
	log.Infof("Body: \n%v", bodyStr)

}
