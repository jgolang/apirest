package apirest

func validate(username, password string) bool {
	if username == "test" && password == "test" {
		return true
	}
	return false
}
