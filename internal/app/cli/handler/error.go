package handler

import (
	"fmt"

	"github.com/bastean/godo/internal/pkg/service/errors"
	"github.com/bastean/godo/internal/pkg/service/record/log"
)

func FormatMessage(what string, who error) string {
	if who != nil {
		return fmt.Sprintf("%s: Error [%s]", what, who.Error())
	}

	return what
}

func ExitByError(err error) {
	var (
		errInvalidValue *errors.ErrInvalidValue
		errFailure      *errors.ErrFailure
		errInternal     *errors.ErrInternal
	)

	var message string

	switch {
	case errors.As(err, &errInvalidValue):
		message = FormatMessage(errInvalidValue.What, errInvalidValue.Who)
	case errors.As(err, &errFailure):
		message = FormatMessage(errFailure.What, errFailure.Who)
	case errors.As(err, &errInternal):
		message = FormatMessage(errInternal.What, errInternal.Who)
	case err != nil:
		message = err.Error()
	}

	if message != "" {
		log.Fatal(message)
	}
}
