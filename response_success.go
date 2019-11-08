package apirest

import "net/http"

// NewSuccessResponse ...
func NewSuccessResponse(title, message string, content interface{}) Response {
	return Response{
		Title:   title,
		Message: message,
		Content: content,
	}
}

// SendSuccessResponse ...
func SendSuccessResponse(title, message string, w http.ResponseWriter) {
	success := Success{
		Title:   title,
		Message: message,
	}
	SendResponse(success, w)
}

// SendSuccessContentResponse ...
func SendSuccessContentResponse(title, message string, content interface{}, w http.ResponseWriter) {
	success := Success{
		Title:   title,
		Message: message,
		Content: content,
	}
	SendResponse(success, w)
}
