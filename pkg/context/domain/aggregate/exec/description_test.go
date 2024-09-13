package exec_test

import (
	"testing"

	"github.com/bastean/godo/pkg/context/domain/aggregate/exec"
	"github.com/bastean/godo/pkg/context/domain/errors"
	"github.com/stretchr/testify/suite"
)

type DescriptionTestSuite struct {
	suite.Suite
}

func (suite *DescriptionTestSuite) SetupTest() {}

func (suite *DescriptionTestSuite) TestWithInvalidLength() {
	value, err := exec.DescriptionWithInvalidLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewDescription",
		What:  "Description must not exceed 100 characters",
		Why: errors.Meta{
			"Description": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitDescriptionSuite(t *testing.T) {
	suite.Run(t, new(DescriptionTestSuite))
}
