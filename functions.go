package apirest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/jgolang/log"
)

var (
	// Username doc ...
	Username = "test"
	// Password doc ...
	Password = "test"
)

func validate(username, password string) bool {
	if username == Username && password == Password {
		return true
	}
	return false
}

// GetHeaderValueString doc ...
func GetHeaderValueString(key string, r *http.Request) (string, Response) {
	value, err := api.GetHeaderValueString(key, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting header value!",
			Message: fmt.Sprintf("The %v key header has not been obtained", key),
		}
	}
	return value, nil
}

// GetHeaderValueInt doc ...
func GetHeaderValueInt(key string, r *http.Request) (int, Response) {
	value, err := api.GetHeaderValueInt(key, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting header value type Int!",
			Message: fmt.Sprintf("The %v key header has not been obtained", key),
		}
	}
	return value, nil
}

// GetHeaderValueInt64 doc ...
func GetHeaderValueInt64(key string, r *http.Request) (int64, Response) {
	value, err := api.GetHeaderValueInt64(key, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting header value type Int64!",
			Message: fmt.Sprintf("The %v key header has not been obtained", key),
		}
	}
	return value, nil
}

// GetHeaderValueFloat64 doc ...
func GetHeaderValueFloat64(key string, r *http.Request) (float64, Response) {
	value, err := api.GetHeaderValueFloat64(key, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting header value type Float64!",
			Message: fmt.Sprintf("The %v key header has not been obtained", key),
		}
	}
	return value, nil
}

// GetHeaderValueBool doc ...
func GetHeaderValueBool(key string, r *http.Request) (bool, Response) {
	value, err := api.GetHeaderValueBool(key, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting header value type Bool!",
			Message: fmt.Sprintf("The %v key header has not been obtained", key),
		}
	}
	return value, nil
}

// GetRouteVarValueString ...
func GetRouteVarValueString(urlVarName string, r *http.Request) (string, Response) {
	value, err := api.GetRouteVarValueString(urlVarName, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting route var!",
			Message: fmt.Sprintf("The route var %v has not been obtained", urlVarName),
		}
	}
	return value, nil
}

// GetRouteVarValueInt ...
func GetRouteVarValueInt(urlVarName string, r *http.Request) (int, Response) {
	value, err := api.GetRouteVarValueInt(urlVarName, r)
	if err != nil {
		log.Error(err)
		return 0, Error{
			Title:   "Error getting route var type Int",
			Message: fmt.Sprintf("The route var %v has not been obtained", urlVarName),
		}
	}
	return value, nil
}

// GetRouteVarValueInt64 ...
func GetRouteVarValueInt64(urlVarName string, r *http.Request) (int64, Response) {
	value, err := api.GetRouteVarValueInt64(urlVarName, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting route var type Int64",
			Message: fmt.Sprintf("The route var %v has not been obtained", urlVarName),
		}
	}
	return value, nil
}

// GetRouteVarValueFloat64 ...
func GetRouteVarValueFloat64(urlVarName string, r *http.Request) (float64, Response) {
	value, err := api.GetRouteVarValueFloat64(urlVarName, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting route var type Float64",
			Message: fmt.Sprintf("The route var %v has not been obtained", urlVarName),
		}
	}
	return value, nil
}

// GetRouteVarValueBool ...
func GetRouteVarValueBool(urlVarName string, r *http.Request) (bool, Response) {
	value, err := api.GetRouteVarValueBool(urlVarName, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting route var type Bool",
			Message: fmt.Sprintf("The route var %v has not been obtained", urlVarName),
		}
	}
	return value, nil
}

// GetQueryParamValueString ...
func GetQueryParamValueString(queryParamName string, r *http.Request) (string, Response) {
	value, err := api.GetQueryParamValueString(queryParamName, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting query param!",
			Message: fmt.Sprintf("The query parameter %v has not been obtained", queryParamName),
		}
	}

	return value, nil

}

// GetQueryParamValueInt ...
func GetQueryParamValueInt(queryParamName string, r *http.Request) (int, Response) {
	value, err := api.GetQueryParamValueInt(queryParamName, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting query param type Int!",
			Message: fmt.Sprintf("The query parameter %v has not been obtained", queryParamName),
		}
	}
	return value, nil
}

// GetQueryParamValueInt64 ...
func GetQueryParamValueInt64(queryParamName string, r *http.Request) (int64, Response) {
	value, err := api.GetQueryParamValueInt64(queryParamName, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting query param type Int64!",
			Message: fmt.Sprintf("The query parameter %v has not been obtained", queryParamName),
		}
	}
	return value, nil
}

// GetQueryParamValueFloat64 ...
func GetQueryParamValueFloat64(queryParamName string, r *http.Request) (float64, Response) {
	value, err := api.GetQueryParamValueFloat64(queryParamName, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting query param type Float64!",
			Message: fmt.Sprintf("The query parameter %v has not been obtained", queryParamName),
		}
	}
	return value, nil
}

// GetQueryParamValueBool ...
func GetQueryParamValueBool(queryParamName string, r *http.Request) (bool, Response) {
	value, err := api.GetQueryParamValueBool(queryParamName, r)
	if err != nil {
		log.Error(err)
		return value, Error{
			Title:   "Error getting query param type Bool!",
			Message: fmt.Sprintf("The query parameter %v has not been obtained", queryParamName),
		}
	}
	return value, nil
}

// UnmarshalBody doc ...
func UnmarshalBody(v interface{}, r *http.Request) Response {
	err := api.UnmarshalBody(v, r)
	if err != nil {
		log.Error(err)
		return Error{
			Title:   "Not unmarshal JSON struct!",
			Message: "Error when unmarshal JSON structure",
		}
	}
	return nil
}

// PrintAPIRequest doc ...
func PrintAPIRequest(method, uri, eventID, form string, header http.Header, rawBody []byte) {
	log.Debugf("REQUEST: [%v] %v", method, uri)
	if eventID != "" {
		log.Debugf("EVENT_ID: ", eventID)
	}
	log.Debugf("HEADER: %v", header)
	if rawBody != nil && len(rawBody) != 0 {
		if len(form) > 2000 && os.Getenv("PRINT_FULL_EVENT") == "" {
			log.Debugf("FORM: ", form[:1000], "••• SKIPPED •••", form[:1000])
		} else {
			log.Debugf("FORM: %v", form)
		}
	}
	if rawBody != nil && len(rawBody) != 0 {
		if len(rawBody) > 2000 && os.Getenv("PRINT_FULL_EVENT") == "" {
			log.Debug("BODY: ", string(rawBody[:1000]), " ••• SKIPPED ••• ", string(rawBody[len(rawBody)-1000:]))
		} else {
			log.Debugf("FORM: ", string(rawBody))
		}
	}
}

// PrintAPIResponse doc ...
func PrintAPIResponse(res *httptest.ResponseRecorder) {
	log.Debugf("RESPONSE: %v", res.Result())
	log.Debugf("STATUS CODE: %v %v", res.Code, http.StatusText(res.Code))
	log.Debugf("HEADER: %v", res.Header())
	responseBody := res.Body.Bytes()
	if responseBody != nil && len(responseBody) != 0 {
		if len(responseBody) > 2000 && os.Getenv("PRINT_FULL_EVENT") == "" {
			log.Debugf("BODY: ", string(responseBody[:1000]), " ••• SKIPPED ••• ", string(responseBody[len(responseBody)-1000:]))
		} else {
			log.Debugf("BODY: ", string(responseBody))
		}
	}
}
