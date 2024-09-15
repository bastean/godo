package cmd_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/bastean/godo/pkg/context/domain/entity/exec"
	"github.com/bastean/godo/pkg/context/domain/errors"
	"github.com/bastean/godo/pkg/context/domain/role"
	"github.com/bastean/godo/pkg/context/domain/service"
	"github.com/bastean/godo/pkg/context/infrastructure/runner/cmd"
	"github.com/stretchr/testify/suite"
)

type CmdTestSuite struct {
	suite.Suite
	sut  role.Executer
	file string
}

func (suite *CmdTestSuite) SetupTest() {
	suite.sut = new(cmd.Cmd)
	suite.file = fmt.Sprintf("temp-%s.tmp", service.Create.UUID())
}

func (suite *CmdTestSuite) TestExecute() {
	task := &exec.Task{
		Commands: []*exec.Command{
			{
				Name: "touch",
				Args: []string{suite.file},
			},
		},
	}

	suite.NoError(suite.sut.Execute(task.Commands[0]))

	result := suite.FileExists(suite.file)

	suite.True(result)
}

func (suite *CmdTestSuite) TestExecuteErrFailure() {
	task := &exec.Task{
		Commands: []*exec.Command{
			{
				Name: "touch",
				Args: []string{
					"-x",
					suite.file,
				},
			},
		},
	}

	err := suite.sut.Execute(task.Commands[0])

	var actual *errors.ErrFailure

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrFailure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Execute",
		What:  actual.What,
		Why: errors.Meta{
			"Name": "touch",
			"Args": []string{"-x", suite.file},
		},
		Who: actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *CmdTestSuite) TearDownTest() {
	os.Remove(suite.file)
}

func TestIntegrationCmdSuite(t *testing.T) {
	suite.Run(t, new(CmdTestSuite))
}
