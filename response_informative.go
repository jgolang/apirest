package apirest

import "net/http"

// Informative info response type the value is "info"
type Informative struct {
	Title      string
	Message    string
	StatusCode int
	Action     string
	Content    interface{}
}

// SetResponse informative ...
func (info Informative) setResponse() (response ResponseData) {
	response = ResponseData{
		Title:      info.Title,
		Message:    info.Message,
		StatusCode: 200,
		Type:       InformativeType,
		Action:     info.Action,
		Content:    info.Content,
	}
	if info.StatusCode != 0 {
		response.StatusCode = info.StatusCode
	}
	return
}

// InformativeResponse ...
func InformativeResponse(title, message string, w http.ResponseWriter) {
	info := Informative{
		Title:   title,
		Message: message,
	}
	SendResponse(info, w)
}
