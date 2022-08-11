package main

import (
	"github.com/chefaku/gorm-restapi/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.RoutesHandler(app)

	app.Listen(":3000")
}
