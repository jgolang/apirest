package apirest

import (
	"net/http"

	"github.com/jgolang/apirest/core"
)

// RequestValidator implementation
type RequestValidator struct{}

// ValidateRequest doc
func (v RequestValidator) ValidateRequest(r *http.Request) (*core.RequestData, error) {
	var request JSONRequest
	err := api.UnmarshalBody(&request, r)
	if err != nil {
		return nil, err
	}
	requestData := core.RequestData{
		DeviceUUID:  request.Info.DeviceUUID,
		DeviceType:  request.Info.DeviceType,
		DeviceOS:    request.Info.DeviceOS,
		OSVersion:   request.Info.OSVersion,
		OSTimezone:  request.Info.OSTimezone,
		AppLanguage: request.Info.AppLanguage,
		AppVersion:  request.Info.AppVersion,
		AppName:     request.Info.AppName,
		SessionID:   request.Info.SessionID,
		Data:        request.Content,
	}
	return &requestData, nil
}
