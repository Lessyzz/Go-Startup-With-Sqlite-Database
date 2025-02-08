package routes

import (
	"fiber-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupIndexRoute(app *fiber.App) {
	route := app.Group("/")
	route.Get("/", handlers.Index_GET)
}
