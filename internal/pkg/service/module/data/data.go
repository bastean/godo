package data

import (
	"github.com/bastean/godo/pkg/application/data/decode"
	"github.com/bastean/godo/pkg/application/data/read"
	"github.com/bastean/godo/pkg/domain/role"
)

type (
	Decoder role.Decoder
	Reader  role.Reader
)

var (
	DecodeJSON *decode.Decode
	ReadLocal  *read.Read
	ReadRemote *read.Read
)

func Start(decoder Decoder, local, remote Reader) {
	DecodeJSON = &decode.Decode{
		Decoder: decoder,
	}

	ReadLocal = &read.Read{
		Reader: local,
	}

	ReadRemote = &read.Read{
		Reader: remote,
	}
}
