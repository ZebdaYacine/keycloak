package middleware

import (
	"echifa/internal/auth"
	"echifa/internal/config"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Keycloak(cfg config.Config, jwks keyfunc.Keyfunc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString, err := extractBearerToken(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		claims := &auth.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, jwks.Keyfunc)
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		if claims.Issuer != cfg.KeycloakIssuer {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid issuer",
			})
		}

		if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token expired",
			})
		}

		println(claims)

		c.Locals("user", claims)

		return c.Next()
	}
}

func extractBearerToken(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid Authorization format")
	}

	return token, nil
}
