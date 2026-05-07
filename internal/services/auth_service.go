package services

import (
	dto "barber_shop/internal/dtos"
	models "barber_shop/internal/model"
	"barber_shop/internal/repositories"
	"barber_shop/internal/utils"
)

func CreateUser(data dto.CreateUserDTO) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(data.Password)

	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashedPassword,
	}

	return repositories.CreateUser(user)
}
