package service

import (
	"net/http"

	"github.com/usmonzodasomon/test-task/internal/client"
	"github.com/usmonzodasomon/test-task/internal/models"
	"github.com/usmonzodasomon/test-task/internal/repository"
)

type Person interface {
	AddPerson(user models.AddPersonInput) (int64, error)
	DeletePerson(id int64) error
}

type Service struct {
	Client
	Person
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Person: NewPersonService(*repos, client.NewUsersClient(&http.Client{})),
	}
}
