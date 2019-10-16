package apirest

import (
	"github.com/jgolang/log"
)

// Init doc ...
func Init() {
	log.ChangeCallerSkip(2)
}
