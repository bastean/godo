package read

import (
	"github.com/bastean/godo/pkg/domain/errors"
	"github.com/bastean/godo/pkg/domain/role"
)

type Read struct {
	role.Reader
}

func (use *Read) Run(route string) ([]byte, error) {
	data, err := use.Reader.Read(route)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return data, nil
}
