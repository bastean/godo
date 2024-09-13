package exec

import (
	"github.com/bastean/godo/pkg/context/domain/errors"
)

func Random() *Exec {
	title := TitleWithValidValue()
	description := DescriptionWithValidValue()
	command := CommandWithValidValue()

	exec, err := New(&Primitive{
		Title:       title.Value,
		Description: description.Value,
		Commands:    []string{command.Value},
	})

	if err != nil {
		errors.Panic(err.Error(), "Random")
	}

	return exec
}

func RandomPrimitive() *Primitive {
	return Random().ToPrimitive()
}
