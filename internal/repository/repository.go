package repository

import (
	"github.com/usmonzodasomon/test-task/internal/models"
	"gorm.io/gorm"
)

type Person interface {
	GetPerson(params models.GetPersonRequest) ([]models.Person, error)
	AddPerson(person models.Person) (int64, error)
	ChangePerson(id int64, person models.Person) error
	DeletePerson(id int64) error
}

type Repository struct {
	Person
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Person: NewPersonRepo(db),
	}
}
