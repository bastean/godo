package exec_test

import (
	"testing"

	"github.com/bastean/godo/pkg/context/domain/aggregate/exec"
	"github.com/bastean/godo/pkg/context/domain/errors"
	"github.com/stretchr/testify/suite"
)

type TitleTestSuite struct {
	suite.Suite
}

func (suite *TitleTestSuite) SetupTest() {}

func (suite *TitleTestSuite) TestWithInvalidLength() {
	value, err := exec.TitleWithInvalidLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewTitle",
		What:  "Title must not exceed 100 characters",
		Why: errors.Meta{
			"Title": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitTitleSuite(t *testing.T) {
	suite.Run(t, new(TitleTestSuite))
}
