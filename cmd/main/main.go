package main

import (
	"github.com/chefaku/gorm-restapi/pkg/db"
	"github.com/chefaku/gorm-restapi/pkg/env"
	"github.com/chefaku/gorm-restapi/pkg/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Loads Environment Variables
	env.LoadEnvFile()

	db.Connect()

	// FIBER SERVER
	app := fiber.New()

	route.RouteHandler(app)

	app.Listen(":3000")
}
