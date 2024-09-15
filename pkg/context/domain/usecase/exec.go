package usecase

import (
	"github.com/bastean/godo/pkg/context/domain/entity/exec"
)

type (
	ExecDo interface {
		Run(*exec.Command) error
	}
)
