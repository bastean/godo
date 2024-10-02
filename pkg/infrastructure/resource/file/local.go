package file

import (
	"fmt"
	"os"

	"github.com/bastean/godo/pkg/domain/errors"
)

type Local struct{}

func (*Local) Read(path string) ([]byte, error) {
	content, err := os.ReadFile(path)

	switch {
	case errors.Is(err, os.ErrNotExist):
		return nil, errors.NewFailure(&errors.Bubble{
			Where: "Read",
			What:  fmt.Sprintf("File does not exist [%s]", path),
		})
	case err != nil:
		return nil, errors.NewFailure(&errors.Bubble{
			Where: "Read",
			What:  fmt.Sprintf("Failure to read file [%s]", path),
			Who:   err,
		})
	}

	return content, nil
}
