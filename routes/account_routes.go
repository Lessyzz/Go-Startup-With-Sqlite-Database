package routes

import (
	"go-starter/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAccountRoutes(app *fiber.App) {
	accountRoutes := app.Group("/account")

	// GET Requests

	accountRoutes.Get("/register", handlers.Register_GET)
	accountRoutes.Get("/login", handlers.Login_GET)

	// POST Requests

	accountRoutes.Post("/register", handlers.Register_POST)
	accountRoutes.Post("/login", handlers.Login_POST)
}
