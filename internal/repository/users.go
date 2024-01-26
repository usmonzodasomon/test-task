package repository

import "gorm.io/gorm"

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) DeleteUser(id int64) error {
	return r.db.Exec("DELETE FROM users WHERE id = $1", id).Error
}
