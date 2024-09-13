package role

import (
	"github.com/bastean/godo/pkg/context/domain/aggregate/exec"
)

type Executer interface {
	Execute(*exec.Command) error
}
