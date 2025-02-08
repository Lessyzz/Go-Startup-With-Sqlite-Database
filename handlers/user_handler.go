package handlers

import (
	"go-starter/database"
	"go-starter/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GET Requests

// @Tags Account Get
// @Router /account/register [get]
func Register_GET(c *fiber.Ctx) error {
	return c.SendFile("./views/register.html")
}

// @Tags Account Get
// @Router /account/login [get]
func Login_GET(c *fiber.Ctx) error {
	return c.SendFile("./views/login.html")
}

// POST Requests

// @Tags Account Post
// @Param user body models.UserRegisterDTO true "User Information"
// @Router /account/register [post]
func Register_POST(c *fiber.Ctx) error {
	var registerDTO models.UserRegisterDTO

	if err := c.BodyParser(&registerDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var existingUser models.User
	result := database.DB.Where("username = ? OR email = ?", registerDTO.Username, registerDTO.Email).First(&existingUser)

	if result.Error == nil && existingUser.ID != "" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
	} else if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	newUser := models.User{
		Username: registerDTO.Username,
		Email:    registerDTO.Email,
		Password: string(registerDTO.Password),
	}

	database.DB.Create(&newUser)

	return c.JSON(fiber.Map{
		"message":  "Registration successful",
		"username": newUser.Username,
		"email":    newUser.Email,
	})
}

// @Tags Account Post
// @Param user body models.UserLoginDTO true "User Information"
// @Router /account/login [post]
func Login_POST(c *fiber.Ctx) error {
	var loginDTO models.UserLoginDTO

	if err := c.BodyParser(&loginDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	result := database.DB.Where("username = ?", loginDTO.Username).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	if loginDTO.Password != user.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong password"})
	}

	return c.JSON(fiber.Map{
		"message":  "Login successful",
		"username": user.Username,
		"email":    user.Email,
	})
}
