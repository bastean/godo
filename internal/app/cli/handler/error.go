package handler

import (
	"fmt"

	"github.com/bastean/godo/internal/pkg/service/errors"
	"github.com/bastean/godo/internal/pkg/service/record/log"
)

func FormatMessage(what string, who error) string {
	if who != nil {
		return fmt.Sprintf("%s: Error [%s]", what, who)
	}

	return what
}

func Error(err error, shouldExit bool) {
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

	switch {
	case message != "" && shouldExit:
		log.Fatal(message)
	case message != "" && !shouldExit:
		log.Error(message)
	case shouldExit:
		log.Fatal("Exited without error message defined")
	}
}

func ExitByError(err error) {
	Error(err, true)
}
