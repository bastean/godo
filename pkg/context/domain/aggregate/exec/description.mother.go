package exec

import (
	"github.com/bastean/godo/pkg/context/domain/errors"
	"github.com/bastean/godo/pkg/context/domain/service"
)

func DescriptionWithValidValue() *Description {
	value, err := NewDescription(service.Create.Regex(`^[A-Za-z0-9\s]{0,100}$`))

	if err != nil {
		errors.Panic(err.Error(), "DescriptionWithValidValue")
	}

	return value
}

func DescriptionWithInvalidLength() (string, error) {
	value := service.Create.Regex(`^[A-Za-z0-9\s]{101,102}$`)

	_, err := NewDescription(value)

	return value, err
}
