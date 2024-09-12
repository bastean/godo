package repository

import (
	"github.com/bastean/godo/pkg/context/user/domain/aggregate/user"
)

type SearchCriteria struct {
	*user.Id
	*user.Email
}

type Repository interface {
	Create(*user.User) error
	Verify(*user.Id) error
	Update(*user.User) error
	Delete(*user.Id) error
	Search(*SearchCriteria) (*user.User, error)
}
