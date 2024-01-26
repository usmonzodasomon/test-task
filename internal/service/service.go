package service

import "github.com/usmonzodasomon/test-task/internal/repository"

type Users interface {
	DeleteUser(id int64) error
}

type Service struct {
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(repos.Users),
	}
}
