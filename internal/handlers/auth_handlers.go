package handlers

import (
	dto "barber_shop/internal/dtos"
	"barber_shop/internal/services"
	"barber_shop/internal/utils"

	"github.com/gofiber/fiber/v3"
)

func Register(c fiber.Ctx) error {
	var body dto.CreateUserDTO

	if err := c.Bind().Body(&body); err != nil {
		return utils.SendError(c, 400, "Invalid body", nil)
	}

	user, err := services.CreateUser(body)

	if err != nil {
		return utils.SendError(c, 500, err.Error(), nil)
	}

	return utils.SendSuccess(c, 201, "User created", user)
}

func Login(c fiber.Ctx) error {

	return nil
}
