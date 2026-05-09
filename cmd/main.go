package main

import (
	"go_store/cmd/router"
	"go_store/configs/database"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

var x = 0
var number int

func main() {
	app := fiber.New()

	database.Connect()

	router.Setup(app)

	app.Use("/", static.New("./public"))

	app.Listen(":3000")
}
