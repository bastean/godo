package cli

import (
	"flag"
	"fmt"

	"github.com/bastean/godo/internal/pkg/service/errors"
	"github.com/bastean/godo/internal/pkg/service/logger/log"
)

const (
	cli = "godo"
)

var (
	task string
)

func usage() {
	log.Logo()

	fmt.Print("Predefined task runner.\n\n")

	fmt.Printf("Usage: %s [flags]\n\n", cli)

	flag.PrintDefaults()
}

func Up() error {
	flag.StringVar(&task, "task", "", "Name of the task to run (required)")

	flag.Usage = usage

	flag.Parse()

	if task == "" {
		return errors.NewInternal(&errors.Bubble{
			Where: "Up",
			What:  "Cannot run an empty task",
		})
	}

	return nil
}
