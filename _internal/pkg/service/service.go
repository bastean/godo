package service

import (
	"context"

	"github.com/bastean/godo/internal/pkg/service/errors"
	"github.com/bastean/godo/internal/pkg/service/module"
)

var (
	err error
)

func Up() error {
	if err = module.Up(); err != nil {
		return errors.BubbleUp(err, "Up")
	}

	return nil
}

func Down(ctx context.Context) error {
	return nil
}
