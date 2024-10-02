package exec

import (
	"github.com/bastean/godo/pkg/application/exec/do"
	"github.com/bastean/godo/pkg/domain/entity/exec"
	"github.com/bastean/godo/pkg/domain/role"
)

type (
	Executer role.Executer
)

type (
	Command = exec.Command
	Task    = exec.Task
	Exec    = exec.Exec
)

var (
	Do *do.Do
)

func Start(executer Executer) {
	Do = &do.Do{
		Executer: executer,
	}
}
