package help

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/bastean/godo/internal/pkg/service/record/log"
)

func Show(cmd *cobra.Command) {
	log.Logo()
	cmd.Help()
	os.Exit(0)
}
