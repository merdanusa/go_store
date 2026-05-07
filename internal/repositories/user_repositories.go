package repositories

import (
	"barber_shop/configs/database"
	models "barber_shop/internal/model"
)

func CreateUser(user models.User) (*models.User, error) {
	result := database.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
