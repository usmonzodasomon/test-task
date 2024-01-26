package db

import (
	"github.com/usmonzodasomon/test-task/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Person{}); err != nil {
		return err
	}
	return nil
}
