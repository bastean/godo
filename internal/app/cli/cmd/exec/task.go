package exec

import (
	"fmt"
	"strings"

	"github.com/bastean/godo/internal/app/cli/util/out"
	"github.com/bastean/godo/internal/pkg/service/errors"
	"github.com/bastean/godo/internal/pkg/service/module/exec"
	"github.com/bastean/godo/internal/pkg/service/record/log"
)

func Do(task *exec.Task) error {
	var (
		err     error
		message string
	)

	totalCommands := len(task.Commands)

	for i, command := range task.Commands {
		message = fmt.Sprintf("(%d/%d) %s %s", i+1, totalCommands, command.Name, strings.Join(command.Args, " "))

		if err = exec.Do.Run(command); err != nil {
			out.Print(log.Error, task.Error, message)
			return errors.BubbleUp(err, "Do")
		}

		log.Success(message)
	}

	return nil
}
