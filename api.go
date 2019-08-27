package apigo

import (
	"github.com/josuegiron/log"
)

func Init() {
	log.ChangeCallerSkip(2)
}
