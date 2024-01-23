package repository

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{}
}
