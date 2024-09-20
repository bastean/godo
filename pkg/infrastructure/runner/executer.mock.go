package runner

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/godo/pkg/domain/entity/exec"
)

type ExecuterMock struct {
	mock.Mock
}

func (executer *ExecuterMock) Execute(command *exec.Command) error {
	executer.Called(command)
	return nil
}
