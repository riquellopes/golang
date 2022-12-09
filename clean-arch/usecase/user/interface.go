package user

import (
	"github.com/riquellopes/golang/clean-arch/entity"
)

type Reader interface {
	List() []*entity.User
}

type Write interface {
	Create(e *entity.User) error
}

type Repository interface {
	Reader
	Write
}

type UserCase interface {
	ListUsers() []*entity.User
}
