package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func CloseDbConnection(dbConn *gorm.DB) error {
	db, err := dbConn.DB()

	if err != nil {
		return fmt.Errorf("error occured on database connection closing: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		return fmt.Errorf("error occured on database connection closing: %s", err.Error())
	}
	return nil
}
