package database

import (
	models "barber_shop/internal/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set, bro where's your env?")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	DB = db

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	log.Println("PostgreSQL connected, we're so back")
}
