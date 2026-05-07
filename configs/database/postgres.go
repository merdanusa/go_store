package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"myapp/internal/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=1234 dbname=myapp port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed")
	}

	DB = db
	
	DB.AutoMigrate(&models.User{})

	log.Println("PostgreSQL connected")
}
