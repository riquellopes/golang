package user

import "github.com/riquellopes/golang/clean-arch/entity"

type inmemory struct {
	m map[int]*entity.User
}

func NewRepository() *inmemory {
	var memo = map[int]*entity.User{}
	return &inmemory{
		m: memo,
	}
}

func (i *inmemory) List() []*entity.User {
	var d []*entity.User
	for _, j := range i.m {
		d = append(d, j)
	}
	return d
}

func (i *inmemory) Create(u *entity.User) error {
	i.m[0] = u
	return nil
}
