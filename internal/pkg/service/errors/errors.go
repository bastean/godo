package errors

import (
	"github.com/bastean/godo/pkg/domain/errors"
)

type (
	Bubble = errors.Bubble
	Meta   = errors.Meta
)

type (
	ErrInvalidValue = errors.ErrInvalidValue
	ErrFailure      = errors.ErrFailure
	ErrInternal     = errors.ErrInternal
)

var (
	NewFailure  = errors.NewFailure
	NewInternal = errors.NewInternal
)

var (
	Panic    = errors.Panic
	BubbleUp = errors.BubbleUp
)

var (
	As = errors.As
	Is = errors.Is
)

func IsNot(err error, target error) bool {
	return err != nil && !Is(err, target)
}
