package service

import (
	"github.com/bastean/godo/internal/pkg/service/module"
)

func Up() {
	module.Start()
}
