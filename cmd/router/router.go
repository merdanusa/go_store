package router

import (
	"go_store/internal/user"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	user.UserRoutes(api)
}
