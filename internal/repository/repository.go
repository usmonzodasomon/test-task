package repository

import "gorm.io/gorm"

type Users interface {
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
