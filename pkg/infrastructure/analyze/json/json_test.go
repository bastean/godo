package json_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/godo/pkg/domain/errors"
	"github.com/bastean/godo/pkg/domain/role"
	"github.com/bastean/godo/pkg/infrastructure/analyze/json"
)

type (
	Path = string
)

type (
	Model struct {
		Ok bool
	}
	Local struct {
		Success Path
		Failure Path
	}
	Data struct {
		*Model
		*Local
	}
)

type JSONTestSuite struct {
	suite.Suite
	sut  role.Decoder
	data *Data
}

func (suite *JSONTestSuite) SetupTest() {
	suite.sut = new(json.JSON)
	suite.data = &Data{
		Local: &Local{
			Success: "testdata/success.json",
			Failure: "testdata/failure.json",
		},
	}
}

func (suite *JSONTestSuite) TestDecode() {
	data, err := os.ReadFile(suite.data.Success)

	suite.NoError(err)

	var actual *Model

	err = suite.sut.Decode(data, &actual)

	suite.NoError(err)

	expected := &Model{
		Ok: true,
	}

	suite.Equal(expected, actual)
}

func (suite *JSONTestSuite) TestDecodeErrFailure() {
	data, err := os.ReadFile(suite.data.Failure)

	suite.NoError(err)

	var actual *errors.ErrFailure

	err = suite.sut.Decode(data, new(Model))

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrFailure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Decode",
		What:  "Invalid type field [Ok] required type is [bool] and [string] was obtained",
	}}

	suite.EqualError(expected, actual.Error())
}

func TestIntegrationJSONSuite(t *testing.T) {
	suite.Run(t, new(JSONTestSuite))
}
