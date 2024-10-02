package decode_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/godo/pkg/application/data/decode"
	"github.com/bastean/godo/pkg/domain/cases"
	"github.com/bastean/godo/pkg/infrastructure/analyze"
)

type DecodeTestSuite struct {
	suite.Suite
	sut     cases.DataDecode
	decoder *analyze.DecoderMock
}

func (suite *DecodeTestSuite) SetupTest() {
	suite.decoder = new(analyze.DecoderMock)

	suite.sut = &decode.Decode{
		Decoder: suite.decoder,
	}
}

func (suite *DecodeTestSuite) TestDecode() {
	data := []byte{}

	target := &struct{}{}

	suite.decoder.On("Decode", data, target)

	suite.NoError(suite.sut.Run(data, target))

	suite.decoder.AssertExpectations(suite.T())
}

func TestUnitDecodeSuite(t *testing.T) {
	suite.Run(t, new(DecodeTestSuite))
}
