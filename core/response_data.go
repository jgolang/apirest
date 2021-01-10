package core

// ResponseFormatted ...
type ResponseFormatted struct {
	Headers    map[string]string
	StatusCode int
	Data       interface{}
}

// ResponseData doc ..
type ResponseData struct {
	Title         string
	Message       string
	StatusCode    int
	ErrorCode     string
	Action        string
	SessionID     string
	ResponseType  ResponseType
	Headers       map[string]string
	AditionalInfo map[string]string
	Data          interface{}
}

// AddAditionalInfo func ...
func (data *ResponseData) AddAditionalInfo(key, value string) {
	data.AditionalInfo[key] = value
}

// AddHeader doc
func (data *ResponseData) AddHeader(key, value string) {
	data.Headers[key] = value
}

// ResponseType contains all the response types identiers
type ResponseType string
