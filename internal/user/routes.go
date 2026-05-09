package user

import "github.com/gofiber/fiber/v3"

func UserRoutes(api fiber.Router) {
	users := api.Group("/users")

	users.Post("/sign-up", CreateUser)
	users.Post("/sign-in", LoginUser)
}
