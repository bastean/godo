package do_test

import (
	"testing"

	"github.com/bastean/godo/pkg/application/exec/do"
	"github.com/bastean/godo/pkg/domain/entity/exec"
	"github.com/bastean/godo/pkg/domain/usecase"
	"github.com/bastean/godo/pkg/infrastructure/runner"
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
	task := &exec.Task{
		Commands: []*exec.Command{
			{
				Name: "godo",
			},
		},
	}

	suite.executer.On("Execute", task.Commands[0])

	suite.NoError(suite.sut.Run(task.Commands[0]))

	suite.executer.AssertExpectations(suite.T())
}

func TestUnitDoSuite(t *testing.T) {
	suite.Run(t, new(DoTestSuite))
}
