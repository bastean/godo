package module

import (
	"github.com/bastean/godo/internal/pkg/service/module/exec"
	"github.com/bastean/godo/internal/pkg/service/runner/cmd"
)

func Start() {
	exec.Start(cmd.Cmd)
}
