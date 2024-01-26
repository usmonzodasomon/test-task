package repository

import (
	"github.com/usmonzodasomon/test-task/internal/models"
	"gorm.io/gorm"
)

type PersonRepo struct {
	db *gorm.DB
}

func NewPersonRepo(db *gorm.DB) *PersonRepo {
	return &PersonRepo{
		db: db,
	}
}

func (r *PersonRepo) AddPerson(user models.Person) (int64, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *PersonRepo) DeletePerson(id int64) error {
	return r.db.Delete(&models.Person{}, id).Error
}
