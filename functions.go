package apigo

import (
	"github.com/josuegiron/log"
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
		log.Error(err)
		return true
	}
	return false
}
