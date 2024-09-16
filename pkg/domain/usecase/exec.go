package usecase

import (
	"github.com/bastean/godo/pkg/domain/entity/exec"
)

type (
	ExecDo interface {
		Run(*exec.Command) error
	}
)
