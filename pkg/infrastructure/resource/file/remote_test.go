package file_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/godo/pkg/domain/errors"
	"github.com/bastean/godo/pkg/domain/role"
	"github.com/bastean/godo/pkg/infrastructure/resource/file"
)

type RemoteTestSuite struct {
	suite.Suite
	sut  role.Reader
	data *Data
}

func (suite *RemoteTestSuite) SetupTest() {
	suite.sut = new(file.Remote)
	suite.data = &Data{
		Local: &Local{
			Success: "testdata/example.exec.json",
		},
		Remote: &Remote{
			Success: "https://raw.githubusercontent.com/bastean/godo/e9068acd9f7cd0fa89310cadfce9e34c1300ead4/configs/example.exec.json",
		},
	}
}

func (suite *RemoteTestSuite) TestRead() {
	actual, err := suite.sut.Read(suite.data.Remote.Success)

	suite.NoError(err)

	expected, err := os.ReadFile(suite.data.Local.Success)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func (suite *RemoteTestSuite) TestReadErrFailure() {
	_, err := suite.sut.Read("")

	var actual *errors.ErrFailure

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrFailure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Read",
		What:  "Cannot obtain the remote resource from []",
		Who:   actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func TestIntegrationRemoteSuite(t *testing.T) {
	suite.Run(t, new(RemoteTestSuite))
}
