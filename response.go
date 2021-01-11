package apirest

import "net/http"

// Response ...
type Response interface {
	Send(w http.ResponseWriter)
}
