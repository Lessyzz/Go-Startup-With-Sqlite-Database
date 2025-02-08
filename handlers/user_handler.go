package handlers

import (
	"fiber-api/database"
	"fiber-api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GET Requests

func Register_GET(c *fiber.Ctx) error {
	return c.SendFile("./views/register.html")
}

func Login_GET(c *fiber.Ctx) error {
	return c.SendFile("./views/login.html")
}

// POST Requests

func Register_POST(c *fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	var existingUser models.User
	result := database.DB.Where("username = ? OR email = ?", username, email).First(&existingUser)

	if result.Error == nil && existingUser.ID != "" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
	} else if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	newUser := models.User{
		Username: username,
		Email:    email,
		Password: string(password),
	}

	database.DB.Create(&newUser)

	return c.JSON(fiber.Map{
		"message":  "Registration successful",
		"username": newUser.Username,
		"email":    newUser.Email,
	})
}

func Login_POST(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	if user.Password != password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong password"})
	}

	return c.JSON(fiber.Map{
		"message":  "Login successful",
		"username": user.Username,
		"email":    user.Email,
	})
}
