package role

import (
	"github.com/bastean/godo/pkg/context/domain/entity/exec"
)

type Executer interface {
	Execute(*exec.Command) error
}
