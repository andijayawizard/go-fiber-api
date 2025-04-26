package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RequireAPIKey(c *fiber.Ctx) error {
	apiKey := c.Get("X-API-Key")
	if apiKey != "rahasia-123" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Invalid API Key",
		})
	}
	return c.Next()
}
