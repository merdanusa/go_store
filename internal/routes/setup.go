package routes

import (
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	AuthRoutes(api)
}
