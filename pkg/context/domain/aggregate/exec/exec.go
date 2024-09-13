package exec

import (
	"github.com/bastean/godo/pkg/context/domain/errors"
)

type Exec struct {
	*Title
	*Description
	Commands []*Command
}

type Primitive struct {
	Title, Description string
	Commands           []string
}

func create(primitive *Primitive) (*Exec, error) {
	title, errTitle := NewTitle(primitive.Title)
	description, errDescription := NewDescription(primitive.Description)
	commands, errCommands := NewCommands(primitive.Commands)

	if err := errors.Join(errTitle, errDescription, errCommands); err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &Exec{
		Title:       title,
		Description: description,
		Commands:    commands,
	}, nil
}

func (exec *Exec) ToPrimitive() *Primitive {
	commands := []string{}

	for _, command := range exec.Commands {
		commands = append(commands, command.Value)
	}

	return &Primitive{
		Title:       exec.Title.Value,
		Description: exec.Description.Value,
		Commands:    commands,
	}
}

func New(primitive *Primitive) (*Exec, error) {
	exec, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	return exec, nil
}
