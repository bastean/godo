package do

import (
	"github.com/bastean/godo/pkg/domain/entity/exec"
	"github.com/bastean/godo/pkg/domain/errors"
	"github.com/bastean/godo/pkg/domain/role"
)

type Do struct {
	role.Executer
}

func (use *Do) Run(cmd *exec.Command) error {
	if err := use.Executer.Execute(cmd); err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
