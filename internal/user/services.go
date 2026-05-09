package user

import (
	"go_store/internal/models"
	utils "go_store/internal/shared"
	hash "go_store/internal/shared/hash"
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

func LoginUserService(data SignInDTO) (string, error) {
	user, err := GetUserByEmail(data.Email)
	if err != nil {

		return "", err
	}

	return utils.GenerateToken(user)
}
