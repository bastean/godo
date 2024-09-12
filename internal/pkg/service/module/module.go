package module

import (
	"github.com/bastean/godo/internal/pkg/service/logger/log"
)

var Module = &struct {
	User string
}{
	User: log.Module("User"),
}

func Up() error {
	log.Starting(Module.User)

	log.Started(Module.User)

	return nil
}
