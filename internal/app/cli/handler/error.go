package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bastean/godo/internal/pkg/service/errors"
	"github.com/bastean/godo/internal/pkg/service/record/log"
)

func ExitByError(err error) {
	var (
		errInvalidValue *errors.ErrInvalidValue
		errFailure      *errors.ErrFailure
		errInternal     *errors.ErrInternal
	)

	var (
		errJsonUnmarshalType *json.UnmarshalTypeError
	)

	var message string

	switch {
	case errors.As(err, &errInvalidValue):
		message = errInvalidValue.What
	case errors.As(err, &errFailure):
		message = errFailure.What
	case errors.As(err, &errInternal):
		message = errInternal.What
	case errors.As(err, &errJsonUnmarshalType):
		message = fmt.Sprintf("Invalid type field [%s] required type is [%s] and [%s] was obtained", errJsonUnmarshalType.Field, errJsonUnmarshalType.Type, errJsonUnmarshalType.Value)
	case errors.Is(err, os.ErrNotExist):
		fallthrough
	case err != nil:
		message = err.Error()
	}

	if message != "" {
		log.Fatal(message)
	}
}
