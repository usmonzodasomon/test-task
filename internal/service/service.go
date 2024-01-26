package service

import (
	"net/http"

	"github.com/usmonzodasomon/test-task/internal/client"
	"github.com/usmonzodasomon/test-task/internal/models"
	"github.com/usmonzodasomon/test-task/internal/repository"
)

type Users interface {
	CreateUser(user models.CreateUserInput) (int64, error)
	DeleteUser(id int64) error
}

type Service struct {
	Client
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(*repos, client.NewUsersClient(&http.Client{})),
	}
}
