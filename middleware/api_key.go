package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func RequireAPIKey(c *fiber.Ctx) error {
	apiKeyFromEnv := os.Getenv("API_KEY")
	apiKeyFromRequest := c.Get("X-API-Key")

	if apiKeyFromRequest != apiKeyFromEnv || apiKeyFromEnv == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Invalid or missing API Key",
		})
	}

	return c.Next()
}
