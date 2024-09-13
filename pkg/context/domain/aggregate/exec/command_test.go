package exec_test

import (
	"testing"

	"github.com/bastean/godo/pkg/context/domain/aggregate/exec"
	"github.com/bastean/godo/pkg/context/domain/errors"
	"github.com/stretchr/testify/suite"
)

type CommandTestSuite struct {
	suite.Suite
}

func (suite *CommandTestSuite) SetupTest() {}

func (suite *CommandTestSuite) TestWithEmptyValue() {
	value, err := exec.CommandWithEmptyValue()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCommand",
		What:  "Command cannot be empty",
		Why: errors.Meta{
			"Command": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitCommandSuite(t *testing.T) {
	suite.Run(t, new(CommandTestSuite))
}
