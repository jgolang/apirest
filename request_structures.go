package apirest

import (
	"encoding/json"
)

// JSONRequest struct used to parse the request content section
type JSONRequest struct {
	Info    JSONRequestInfo `json:"info,omitempty"`
	Content json.RawMessage `json:"content"`
}

// JSONRequestInfo request info section fields for encrypted requests
type JSONRequestInfo struct {
	DeviceUUID  string `json:"uuid,omitempty"`
	DeviceType  string `json:"device,omitempty"`
	OS          string `json:"os,omitempty"`
	OSVersion   string `json:"os_version,omitempty"`
	OSTimezone  int    `json:"timezone,omitempty"`
	AppLanguage string `json:"lang,omitempty"`
	AppVersion  string `json:"app_version,omitempty"`
	SessionID   string `json:"session_id,omitempty"`
	AppName     string `json:"app_name,omitempty"`
}
