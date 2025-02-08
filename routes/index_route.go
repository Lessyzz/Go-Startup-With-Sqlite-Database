package routes

import (
	"go-starter/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupIndexRoute(app *fiber.App) {
	route := app.Group("/")
	route.Get("/", handlers.Index_GET)
}
