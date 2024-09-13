package exec

import (
	"fmt"
	"strings"

	"github.com/bastean/godo/pkg/context/domain/errors"
	"github.com/bastean/godo/pkg/context/domain/service"
)

const (
	TitleMaxCharactersLength = "100"
)

type Title struct {
	Value string `validate:"lte=100"`
}

func NewTitle(value string) (*Title, error) {
	value = strings.TrimSpace(value)

	valueObj := &Title{
		Value: value,
	}

	if service.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewTitle",
			What:  fmt.Sprintf("Title must not exceed %s characters", TitleMaxCharactersLength),
			Why: errors.Meta{
				"Title": value,
			},
		})
	}

	return valueObj, nil
}
