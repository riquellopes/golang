package user

import "github.com/riquellopes/golang/clean-arch/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) ListUsers() []*entity.User {
	return s.repo.List()
}

func (s *Service) CreateUser(name string, age int) (*entity.User, error) {
	user, err := entity.NewUser(name, age)

	if err != nil {
		return nil, err
	}

	s.repo.Create(user)
	return user, nil
}
