package user

import (
	"go_store/configs/database"
	"go_store/internal/models"
)

func CreateUserRepository(user models.User) (*models.User, error) {
	result := database.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}