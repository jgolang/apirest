package apirest

import (
	"go.uber.org/zap"
)

// Check doc...
func Check(err error) bool {
	if err != nil {
		return true
	}
	return false
}

// Checkp doc...
func Checkp(err error) bool {
	if err != nil {
		logger := zap.S()
		defer logger.Sync()
		logger.Error(err)
		return true
	}
	return false
}
