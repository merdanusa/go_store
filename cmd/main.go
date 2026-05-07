package main

import (
	"barber_shop/configs/database"
	"barber_shop/internal/routes"

	"github.com/gofiber/fiber/v3"
)

var x = 0
var number int

func main() {
	app := fiber.New()

	database.Connect()

	routes.Setup(app)

	app.Listen(":3000")
}
