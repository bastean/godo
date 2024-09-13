package exec

import (
	"github.com/bastean/godo/pkg/context/domain/errors"
	"github.com/bastean/godo/pkg/context/domain/service"
)

type Command struct {
	Value string `validate:"required"`
}

func NewCommand(value string) (*Command, error) {
	valueObj := &Command{
		Value: value,
	}

	if service.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewCommand",
			What:  "Command cannot be empty",
			Why: errors.Meta{
				"Command": value,
			},
		})
	}

	return valueObj, nil
}

func NewCommands(values []string) ([]*Command, error) {
	var (
		err     error
		command *Command
	)

	commands := []*Command{}

	for _, value := range values {
		command, err = NewCommand(value)

		if err != nil {
			return nil, errors.BubbleUp(err, "NewCommands")
		}

		commands = append(commands, command)
	}

	return commands, nil
}
