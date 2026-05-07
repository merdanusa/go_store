package routes

import (
	"barber_shop/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func AuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/logic", handlers.Login)
	auth.Post("/register", handlers.Register)
}
