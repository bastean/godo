package resource

import (
	"github.com/stretchr/testify/mock"
)

type ReaderMock struct {
	mock.Mock
}

func (reader *ReaderMock) Read(route string) ([]byte, error) {
	args := reader.Called(route)
	return args.Get(0).([]byte), nil
}
