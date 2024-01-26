package repository

import (
	"github.com/usmonzodasomon/test-task/internal/models"
	"gorm.io/gorm"
)

type Person interface {
	AddPerson(user models.Person) (int64, error)
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
