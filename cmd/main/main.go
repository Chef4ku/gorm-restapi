package main

import (
	"github.com/chefaku/gorm-restapi/pkg/env"
	"github.com/chefaku/gorm-restapi/pkg/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Environment Variables
	env.LoadEnvFile()

	// FIBER SERVER
	app := fiber.New()

	route.RouteHandler(app)

	app.Listen(":3000")
}
