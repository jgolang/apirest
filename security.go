package apirest

var (
	// Username doc ...
	Username = "test"
	// Password doc ...
	Password = "test"
)

// Security core interface implement
type Security struct {
}

// ValidateBasicToken doc ...
func (s *Security) ValidateBasicToken(token string) bool {

	return true
}

// ValidateBearerToken doc ...
func (s *Security) ValidateBearerToken(token string) bool {

	return true
}

// ValidateCustomToken doc ...
func (s *Security) ValidateCustomToken(validator func(string) bool) bool {
	return validator("")
}

func validate(username, password string) bool {
	if username == Username && password == Password {
		return true
	}
	return false
}
