package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/bastean/godo/internal/app/cli/cmd/exec"
	"github.com/bastean/godo/internal/app/cli/service/help"
)

const (
	cli = "godo"
)

var cmd = &cobra.Command{
	Use:   cli,
	Short: "Predefined task runner.",
	Run: func(cmd *cobra.Command, args []string) {
		help.Show(cmd)
	},
}

func init() {
	cmd.AddCommand(exec.Command)
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
