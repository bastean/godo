package module

import (
	"github.com/bastean/godo/internal/pkg/service/analyze/json"
	"github.com/bastean/godo/internal/pkg/service/module/data"
	"github.com/bastean/godo/internal/pkg/service/module/exec"
	"github.com/bastean/godo/internal/pkg/service/resource/file"
	"github.com/bastean/godo/internal/pkg/service/run/cmd"
)

func Start() {
	exec.Start(cmd.CMD)
	data.Start(json.JSON, file.Local, file.Remote)
}
