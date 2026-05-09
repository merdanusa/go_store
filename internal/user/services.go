package user

import (
	"barber_shop/internal/models"
	hash "barber_shop/internal/shared/hash"
)

func CreateUserService(data SignUpDTO) (*models.User, error) {
	hashed, err := hash.HashPassword(data.Password)

	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashed,
	}

	return CreateUserRepository(user)
}
