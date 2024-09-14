package usecase

import (
	"github.com/bastean/godo/pkg/context/domain/aggregate/exec"
)

type (
	ExecDo interface {
		Run(exec.Command) error
	}
)
