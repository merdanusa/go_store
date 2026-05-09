package user

import (
	"barber_shop/configs/database"
	"barber_shop/internal/models"
)

func CreateUserRepository(user models.User) (*models.User, error) {
	result := database.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}