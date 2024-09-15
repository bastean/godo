package runner

import (
	"github.com/bastean/godo/pkg/context/domain/entity/exec"
	"github.com/stretchr/testify/mock"
)

type ExecuterMock struct {
	mock.Mock
}

func (executer *ExecuterMock) Execute(command *exec.Command) error {
	executer.Called(command)
	return nil
}
