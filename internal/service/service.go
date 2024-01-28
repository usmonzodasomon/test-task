package service

import (
	"net/http"

	"github.com/usmonzodasomon/test-task/internal/client"
	"github.com/usmonzodasomon/test-task/internal/models"
	"github.com/usmonzodasomon/test-task/internal/repository"
)

type Person interface {
	GetPerson(params models.GetPersonRequest) ([]models.Person, error)
	AddPerson(person models.AddPersonInput) (int64, error)
	ChangePerson(id int64, person models.Person) error
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
