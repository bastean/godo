package do_test

import (
	"testing"

	"github.com/bastean/godo/pkg/context/application/exec/do"
	"github.com/bastean/godo/pkg/context/domain/aggregate/exec"
	"github.com/bastean/godo/pkg/context/domain/usecase"
	"github.com/bastean/godo/pkg/context/infrastructure/runner"
	"github.com/stretchr/testify/suite"
)

type DoTestSuite struct {
	suite.Suite
	sut      usecase.ExecDo
	executer *runner.ExecuterMock
}

func (suite *DoTestSuite) SetupTest() {
	suite.executer = new(runner.ExecuterMock)

	suite.sut = &do.Do{
		Executer: suite.executer,
	}
}

func (suite *DoTestSuite) TestDo() {
	task := &exec.Exec{
		Commands: []string{
			"godo",
		},
	}

	suite.executer.On("Execute", task.Commands[0])

	suite.NoError(suite.sut.Run(task.Commands[0]))

	suite.executer.AssertExpectations(suite.T())
}

func TestUnitDoSuite(t *testing.T) {
	suite.Run(t, new(DoTestSuite))
}
