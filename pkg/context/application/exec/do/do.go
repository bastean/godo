package do

import (
	"github.com/bastean/godo/pkg/context/domain/aggregate/exec"
	"github.com/bastean/godo/pkg/context/domain/errors"
	"github.com/bastean/godo/pkg/context/domain/role"
	"github.com/bastean/godo/pkg/context/domain/usecase"
)

type Do struct {
	role.Logger
	role.Executer
}

func (use *Do) Run(primitive *exec.Primitive) error {
	task, err := exec.New(primitive)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	use.Logger.Info(task.Title.Value)

	use.Logger.Info(task.Description.Value)

	for _, command := range task.Commands {
		if err = use.Executer.Execute(command); err != nil {
			return errors.BubbleUp(err, "Run")
		}
	}

	return nil
}

func New(logger role.Logger, executer role.Executer) usecase.ExecDo {
	return &Do{
		Logger:   logger,
		Executer: executer,
	}
}
