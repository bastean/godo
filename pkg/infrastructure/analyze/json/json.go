package json

import (
	"encoding/json"
	"fmt"

	"github.com/bastean/godo/pkg/domain/errors"
)

type JSON struct{}

func (*JSON) Decode(data []byte, target any) error {
	var errJsonUnmarshalType *json.UnmarshalTypeError

	err := json.Unmarshal(data, &target)

	switch {
	case errors.As(err, &errJsonUnmarshalType):
		return errors.NewFailure(&errors.Bubble{
			Where: "Decode",
			What:  fmt.Sprintf("Invalid type field [%s] required type is [%s] and [%s] was obtained", errJsonUnmarshalType.Field, errJsonUnmarshalType.Type, errJsonUnmarshalType.Value),
		})
	case err != nil:
		return errors.NewFailure(&errors.Bubble{
			Where: "Decode",
			What:  "Failure to read JSON",
			Who:   err,
		})
	}

	return nil
}
