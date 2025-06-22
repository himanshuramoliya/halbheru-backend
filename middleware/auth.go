package middleware

import (
	"strings"

	"halbheru-backend/utils"

	"github.com/gofiber/fiber/v2"
)

// AuthRequired middleware checks for valid JWT token
func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Authorization header required",
			})
		}

		// Check if it starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid authorization format",
			})
		}

		// Extract token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Token required",
			})
		}

		// Validate token
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Store user ID in context for use in handlers
		c.Locals("userID", claims.UserID)

		return c.Next()
	}
}
