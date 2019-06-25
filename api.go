package apigolang

import (
	"github.com/josuegiron/log"
)

func init() {
	log.ChangeCallerSkip(2)
}
