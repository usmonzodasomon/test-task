package service

import "github.com/usmonzodasomon/test-task/internal/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
