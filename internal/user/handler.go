package user

import (
	utils "go_store/internal/shared"

	"github.com/gofiber/fiber/v3"
)

func CreateUser(c fiber.Ctx) error {
	var body SignUpDTO

	if err := c.Bind().Body(&body); err != nil {
		return utils.SendError(c, 400, "Invalid body", nil)
	}

	user, err := CreateUserService(body)

	if err != nil {
		return utils.SendError(c, 500, err.Error(), nil)
	}

	return c.JSON(user)
}

func LoginUser(c fiber.Ctx) error {
	var body SignInDTO

	if err := c.Bind().Body(&body); err != nil {
		return utils.SendError(c, 400, "Invalid body", nil)
	}

	token, err := LoginUserService(body)
	if err != nil {
		return utils.SendError(c, 500, err.Error(), nil)
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
