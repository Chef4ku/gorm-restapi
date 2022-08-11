package routes

import "github.com/gofiber/fiber/v2"

func index(app *fiber.App) {
	app.Static("/", "./client/dist")
}
