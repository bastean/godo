package cmd

import (
	ex "os/exec"
	"strings"

	"github.com/bastean/godo/pkg/context/domain/aggregate/exec"
	"github.com/bastean/godo/pkg/context/domain/errors"
)

type Cmd struct{}

func (*Cmd) Execute(command *exec.Command) error {
	values := strings.Split(command.Value, " ")

	name := values[0]
	args := values[1:]

	if output, err := ex.Command(name, args...).CombinedOutput(); err != nil {
		return errors.NewFailure(&errors.Bubble{
			Where: "Execute",
			What:  "Cannot execute command",
			Why: errors.Meta{
				"Name":   name,
				"Args":   args,
				"Output": string(output),
			},
			Who: err,
		})
	}

	return nil
}

func New() *Cmd {
	return new(Cmd)
}
