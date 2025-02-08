package main

import (
	"fmt"
	"go-starter/database"
	"go-starter/models"
	"go-starter/routes"
	"go-starter/tokenization"

	"github.com/gofiber/fiber/v2"

	_ "go-starter/docs"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupSwagger(app *fiber.App) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}

func SetupRoutes(app *fiber.App) {
	routes.SetupIndexRoute(app)
	routes.SetupAccountRoutes(app)
}

func JWTStuffs() {
	token, _ := tokenization.GenerateJWT("test")

	_, claims, err := tokenization.ValidateJWT(token)

	if err != nil {
		fmt.Println("Token error:", err)
	} else {
		fmt.Println("Username:", claims.Username)
	}
}

func main() {
	app := fiber.New()
	app.Static("/", "./public")

	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.User{})

	SetupSwagger(app) // Should be disable on production
	SetupRoutes(app)

	app.Listen(":3000")
}
