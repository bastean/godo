package exec

import (
	"github.com/bastean/godo/pkg/context/domain/errors"
	"github.com/bastean/godo/pkg/context/domain/service"
)

func TitleWithValidValue() *Title {
	value, err := NewTitle(service.Create.Regex(`^[A-Za-z0-9\s]{0,100}$`))

	if err != nil {
		errors.Panic(err.Error(), "TitleWithValidValue")
	}

	return value
}

func TitleWithInvalidLength() (string, error) {
	value := service.Create.Regex(`^[A-Za-z0-9\s]{101,102}$`)

	_, err := NewTitle(value)

	return value, err
}
