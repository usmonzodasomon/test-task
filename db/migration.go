package db

import (
	"github.com/usmonzodasomon/test-task/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	return nil
}
