package read_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/godo/pkg/application/data/read"
	"github.com/bastean/godo/pkg/domain/cases"
	"github.com/bastean/godo/pkg/infrastructure/resource"
)

type ReadTestSuite struct {
	suite.Suite
	sut    cases.DataRead
	reader *resource.ReaderMock
}

func (suite *ReadTestSuite) SetupTest() {
	suite.reader = new(resource.ReaderMock)

	suite.sut = &read.Read{
		Reader: suite.reader,
	}
}

func (suite *ReadTestSuite) TestRead() {
	route := "success/"

	expected := []byte("success")

	suite.reader.On("Read", route).Return(expected)

	actual, err := suite.sut.Run(route)

	suite.NoError(err)

	suite.reader.AssertExpectations(suite.T())

	suite.Equal(expected, actual)
}

func TestUnitReadSuite(t *testing.T) {
	suite.Run(t, new(ReadTestSuite))
}
