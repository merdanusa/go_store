package database

import (
	"errors"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go_store/internal/models"
)

var DB *gorm.DB

func Connect() error {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		return errors.New("DATABASE_URL is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	log.Println("Running migrations...")

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	log.Println("Database ready")
	return nil
}
