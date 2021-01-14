package apirest

import (
	"encoding/base64"
	"strings"
)

var (
	// Username doc ...
	Username = "test"
	// Password doc ...
	Password = "test"
)

// Security core interface implement
type Security struct {
}

// CustomTokenValidatorFunc doc ...
var CustomTokenValidatorFunc func(string) bool

// ValidateBasicToken doc ...
func (s *Security) ValidateBasicToken(token string) bool {
	payload, _ := base64.StdEncoding.DecodeString(token)
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 || !validate(pair[0], pair[1]) {
		return false
	}
	return true
}

// ValidateBearerToken doc ...
func (s *Security) ValidateBearerToken(token string) bool {

	return true
}

// ValidateCustomToken doc ...
func (s *Security) ValidateCustomToken(token string, validator func(string) bool) bool {
	return validator(token)
}

func validate(username, password string) bool {
	if username == Username && password == Password {
		return true
	}
	return false
}

func validateCustomToken(token string) bool {
	return api.ValidateCustomToken(token, CustomTokenValidatorFunc)
}
