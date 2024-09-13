package exec

import (
	"fmt"
	"strings"

	"github.com/bastean/godo/pkg/context/domain/errors"
	"github.com/bastean/godo/pkg/context/domain/service"
)

const (
	DescriptionMaxCharactersLength = "100"
)

type Description struct {
	Value string `validate:"lte=100"`
}

func NewDescription(value string) (*Description, error) {
	value = strings.TrimSpace(value)

	valueObj := &Description{
		Value: value,
	}

	if service.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewDescription",
			What:  fmt.Sprintf("Description must not exceed %s characters", DescriptionMaxCharactersLength),
			Why: errors.Meta{
				"Description": value,
			},
		})
	}

	return valueObj, nil
}
