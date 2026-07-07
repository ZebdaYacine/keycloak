package main

import (
	"context"
	"echifa/internal/config"
	"echifa/internal/handlers"
	"echifa/internal/middleware"
	"log"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	jwks, err := keyfunc.NewDefaultCtx(context.Background(), []string{cfg.KeycloakJWKS})
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", handlers.Health)

	api := app.Group("/api", middleware.Keycloak(cfg, jwks))

	api.Get("/private", handlers.Private)

	api.Get("/admin/cards", middleware.RequirePermission(middleware.CardAdmin), handlers.Cards)
	// api.Get("/user/cards", middleware.RequirePermission(middleware.CardUser), handlers.Cards)

	log.Fatal(app.Listen(":" + cfg.Port))
}
