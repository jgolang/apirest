package apirest

// Check doc...
func Check(err error) bool {
	if err != nil {
		return true
	}
	return false
}

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
