package exec

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"

	"github.com/bastean/godo/internal/app/cli/handler"
	"github.com/bastean/godo/internal/app/cli/util/help"
	"github.com/bastean/godo/internal/app/cli/util/out"
	"github.com/bastean/godo/internal/pkg/service/module/exec"
	"github.com/bastean/godo/internal/pkg/service/record/log"
)

var Flag = &struct {
	Config, ConfigShort string
}{
	Config:      "config",
	ConfigShort: "c",
}

var Command = &cobra.Command{
	Use:   "exec",
	Short: "Execute a list of tasks from a file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			help.Show(cmd)
		}

		filepath, err := cmd.Flags().GetString(Flag.Config)

		if err != nil {
			handler.ExitByError(err)
		}

		list, err := os.ReadFile(filepath)

		if err != nil {
			handler.ExitByError(err)
		}

		tasks := &struct {
			Tasks []*exec.Task
		}{}

		err = json.Unmarshal(list, &tasks)

		if err != nil {
			handler.ExitByError(err)
		}

		for _, task := range tasks.Tasks {
			out.Print(log.Info, task.Title, task.Description)

			if err = Do(task); err != nil {
				handler.ExitByError(err)
			}

			out.Print(log.Success, task.Success)
		}
	},
}

func init() {
	Command.Flags().StringP(Flag.Config, Flag.ConfigShort, "", "Configuration filepath (required)")
	Command.MarkFlagRequired(Flag.Config)
}
