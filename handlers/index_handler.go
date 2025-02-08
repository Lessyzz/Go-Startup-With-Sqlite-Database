package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GET Requests

// @Tags Index Get
// @Router / [get]
func Index_GET(c *fiber.Ctx) error {
	return c.SendFile("./views/index.html")
}
