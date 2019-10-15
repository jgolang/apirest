package apigo

import (
	"github.com/josuegiron/log"
)

// Init doc ...
func Init() {
	log.ChangeCallerSkip(2)
}
