package decode

import (
	"github.com/bastean/godo/pkg/domain/errors"
	"github.com/bastean/godo/pkg/domain/role"
)

type Decode struct {
	role.Decoder
}

func (use *Decode) Run(data []byte, target any) error {
	if err := use.Decoder.Decode(data, target); err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
