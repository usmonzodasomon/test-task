package service

import "github.com/usmonzodasomon/test-task/internal/repository"

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return NewUsersService(repo)
}

func (s *UsersService) DeleteUser(id int64) error {
	return s.repo.DeleteUser(id)
}
