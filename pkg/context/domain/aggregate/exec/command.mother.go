package exec

import (
	"github.com/bastean/godo/pkg/context/domain/errors"
	"github.com/bastean/godo/pkg/context/domain/service"
)

func CommandWithValidValue() *Command {
	value, err := NewCommand(service.Create.LoremIpsumSentence(service.Create.IntRange(1, 3)))

	if err != nil {
		errors.Panic(err.Error(), "CommandWithValidValue")
	}

	return value
}

func CommandWithEmptyValue() (string, error) {
	value := ""

	_, err := NewCommand(value)

	return value, err
}
