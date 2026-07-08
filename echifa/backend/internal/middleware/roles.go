package middleware

import (
	"echifa/internal/auth"

	"github.com/gofiber/fiber/v2"
)

const (
	CardAdmin = "cards:riud"
	CardUser  = "cards:riu"
)

func RequirePermission(permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("user").(*auth.Claims)
		if !ok || claims == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing user context",
			})
		}

		if claims.HasRealmRole(CardAdmin) || claims.HasRealmRole(permission) {
			return c.Next()
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied",
		})
	}
}
