package read

import (
	"strings"

	"github.com/bastean/godo/internal/pkg/service/module/data"
)

func File(route string) ([]byte, error) {
	if strings.HasPrefix(route, "http") {
		return data.ReadRemote.Run(route)
	}

	return data.ReadLocal.Run(route)
}

func JSON(value []byte, target any) error {
	return data.DecodeJSON.Run(value, &target)
}
