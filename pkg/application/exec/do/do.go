package do

import (
	"github.com/bastean/godo/pkg/domain/entity/exec"
	"github.com/bastean/godo/pkg/domain/errors"
	"github.com/bastean/godo/pkg/domain/role"
)

type Do struct {
	role.Executer
}

func (use *Do) Run(command *exec.Command) error {
	if err := use.Executer.Execute(command); err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
