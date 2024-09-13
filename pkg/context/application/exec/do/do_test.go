package do_test

import (
	"testing"

	"github.com/bastean/godo/pkg/context/application/exec/do"
	"github.com/bastean/godo/pkg/context/domain/aggregate/exec"
	"github.com/bastean/godo/pkg/context/domain/usecase"
	"github.com/bastean/godo/pkg/context/infrastructure/record"
	"github.com/bastean/godo/pkg/context/infrastructure/runner"
	"github.com/stretchr/testify/suite"
)

type DoTestSuite struct {
	suite.Suite
	sut      usecase.ExecDo
	logger   *record.LoggerMock
	executer *runner.ExecuterMock
}

func (suite *DoTestSuite) SetupTest() {
	suite.logger = new(record.LoggerMock)

	suite.executer = new(runner.ExecuterMock)

	suite.sut = do.New(suite.logger, suite.executer)
}

func (suite *DoTestSuite) TestDo() {
	random := exec.Random()

	primitive := random.ToPrimitive()

	suite.logger.On("Info", random.Title.Value)

	suite.logger.On("Info", random.Description.Value)

	suite.executer.On("Execute", random.Commands[0])

	suite.NoError(suite.sut.Run(primitive))

	suite.logger.AssertExpectations(suite.T())

	suite.executer.AssertExpectations(suite.T())
}

func TestUnitDoSuite(t *testing.T) {
	suite.Run(t, new(DoTestSuite))
}
