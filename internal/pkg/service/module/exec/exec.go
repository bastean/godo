package exec

import (
	"github.com/bastean/godo/pkg/context/application/exec/do"
	"github.com/bastean/godo/pkg/context/domain/entity/exec"
	"github.com/bastean/godo/pkg/context/domain/role"
)

type (
	Executer role.Executer
)

type (
	Task    = exec.Task
	Command = exec.Command
)

func Start(executer Executer) {
	Do = &do.Do{
		Executer: executer,
	}
}
