package read

import (
	"github.com/bastean/godo/pkg/context/shared/domain/errors"
	"github.com/bastean/godo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/godo/pkg/context/user/domain/repository"
)

type Read struct {
	repository.Repository
}

func (use *Read) Run(id *user.Id) (*user.User, error) {
	found, err := use.Repository.Search(&repository.SearchCriteria{
		Id: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return found, nil
}
