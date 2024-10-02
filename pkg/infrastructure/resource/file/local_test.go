package file_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/godo/pkg/domain/errors"
	"github.com/bastean/godo/pkg/domain/role"
	"github.com/bastean/godo/pkg/infrastructure/resource/file"
)

type LocalTestSuite struct {
	suite.Suite
	sut  role.Reader
	data *Data
}

func (suite *LocalTestSuite) SetupTest() {
	suite.sut = new(file.Local)
	suite.data = &Data{
		Local: &Local{
			Success: "testdata/success.json",
		},
	}
}

func (suite *LocalTestSuite) TestRead() {
	actual, err := suite.sut.Read(suite.data.Local.Success)

	suite.NoError(err)

	expected, err := os.ReadFile(suite.data.Local.Success)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func (suite *LocalTestSuite) TestReadErrFailure() {
	path := ""

	_, err := suite.sut.Read(path)

	var actual *errors.ErrFailure

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrFailure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Read",
		What:  fmt.Sprintf("File does not exist [%s]", path),
	}}

	suite.EqualError(expected, actual.Error())
}

func TestIntegrationLocalSuite(t *testing.T) {
	suite.Run(t, new(LocalTestSuite))
}
