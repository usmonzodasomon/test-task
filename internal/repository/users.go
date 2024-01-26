package repository

import (
	"github.com/usmonzodasomon/test-task/internal/models"
	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) CreateUser(user models.User) (int64, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *UsersRepo) DeleteUser(id int64) error {
	return r.db.Exec("DELETE FROM users WHERE id = $1", id).Error
}
