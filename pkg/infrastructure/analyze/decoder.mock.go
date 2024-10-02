package analyze

import (
	"github.com/stretchr/testify/mock"
)

type DecoderMock struct {
	mock.Mock
}

func (decoder *DecoderMock) Decode(data []byte, target any) error {
	decoder.Called(data, target)
	return nil
}
