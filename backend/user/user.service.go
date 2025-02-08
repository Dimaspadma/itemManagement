package user

import (
	"awesomeProject/helper"
	"context"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) CreateUser(ctx context.Context, username string, password string) (*User, error) {
	user := &User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now().UTC(),
	}

	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUsers(ctx context.Context) []*User {
	users, err := s.repo.GetUsers(ctx)
	helper.PanicIfError(err)

	return users
}

func (s *Service) DeleteUserById(ctx context.Context, id string) error {
	err := s.repo.DeleteUser(ctx, id)
	return err
}
