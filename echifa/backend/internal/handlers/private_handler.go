package handlers

import (
	"echifa/internal/auth"

	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "E-CHIFA API is running",
	})
}

func Private(c *fiber.Ctx) error {
	claims := c.Locals("user").(*auth.Claims)

	return c.JSON(fiber.Map{
		"message": "Protected data from E-CHIFA API",
		"subject": claims.Subject,
		"issuer":  claims.Issuer,
		"roles":   claims.RealmAccess.Roles,
	})
}

func Admin(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Admin area",
	})
}
