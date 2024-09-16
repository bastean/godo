package cmd

import (
	ex "os/exec"

	"github.com/bastean/godo/pkg/domain/entity/exec"
	"github.com/bastean/godo/pkg/domain/errors"
)

type Cmd struct{}

func (*Cmd) Execute(command *exec.Command) error {
	if output, err := ex.Command(command.Name, command.Args...).CombinedOutput(); err != nil {
		return errors.NewFailure(&errors.Bubble{
			Where: "Execute",
			What:  string(output),
			Why: errors.Meta{
				"Name": command.Name,
				"Args": command.Args,
			},
			Who: err,
		})
	}

	return nil
}
