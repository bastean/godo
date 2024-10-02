package exec

import (
	"fmt"
	"strings"

	"github.com/bastean/godo/internal/pkg/service/errors"
	"github.com/bastean/godo/internal/pkg/service/module/exec"
	"github.com/bastean/godo/internal/pkg/service/record/log"
)

func Do(task *exec.Task) error {
	var (
		err           error
		actualCommand string
	)

	totalCommands := len(task.Commands)

	for i, command := range task.Commands {
		actualCommand = fmt.Sprintf("(%d/%d) %s %s", i+1, totalCommands, command.Name, strings.Join(command.Args, " "))

		if err = exec.Do.Run(command); err != nil {
			log.Error(actualCommand)
			return errors.BubbleUp(err, "Do")
		}

		log.Success(actualCommand)
	}

	return nil
}
