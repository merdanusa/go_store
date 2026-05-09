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

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
