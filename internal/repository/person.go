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

func (r *PersonRepo) GetPersonByID(id int64) (models.Person, error) {
	var person models.Person
	if err := r.db.Where("id = ?", id).First(&person).Error; err != nil {
		return models.Person{}, err
	}
	return person, nil
}

func (r *PersonRepo) GetPerson(params models.GetPersonRequest) ([]models.Person, error) {
	paramsMap := map[string]interface{}{}

	if params.Age != -1 {
		paramsMap["age"] = params.Age
	}

	if params.Gender != "" {
		paramsMap["gender"] = params.Gender
	}

	if params.Nationality != "" {
		paramsMap["nationality"] = params.Nationality
	}

	var people []models.Person
	if err := r.db.Where(paramsMap).Find(&people).Limit(params.Limit).Offset(params.Offset).Error; err != nil {
		return nil, err
	}
	return people, nil
}

func (r *PersonRepo) AddPerson(person models.Person) (int64, error) {
	if err := r.db.Create(&person).Error; err != nil {
		return 0, err
	}
	return person.ID, nil
}

func (r *PersonRepo) ChangePerson(id int64, person models.Person) error {
	updates := map[string]interface{}{}

	if person.Name != "" {
		updates["name"] = person.Name
	}

	if person.Surname != "" {
		updates["surname"] = person.Surname
	}

	if person.Patronomic != "" {
		updates["patronomic"] = person.Patronomic
	}

	if person.Age != 0 {
		updates["age"] = person.Age
	}

	if person.Gender != "" {
		updates["gender"] = person.Gender
	}

	if person.Nationality != "" {
		updates["nationality"] = person.Nationality
	}

	return r.db.Model(&models.Person{}).Where("id = ?", id).Updates(updates).Error
}

func (r *PersonRepo) DeletePerson(id int64) error {
	return r.db.Delete(&models.Person{}, id).Error
}
