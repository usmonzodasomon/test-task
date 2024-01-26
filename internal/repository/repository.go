package repository

import (
	"github.com/usmonzodasomon/test-task/internal/models"
	"gorm.io/gorm"
)

type Users interface {
	CreateUser(user models.User) (int64, error)
	DeleteUser(id int64) error
}

type Repository struct {
	Users
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users: NewUsersRepo(db),
	}
}
