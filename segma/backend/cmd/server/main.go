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
	api.Get("/admin", middleware.RequirePermission("admin"), handlers.Admin)


	api.Get(
		"/admin/segma",
		middleware.RequirePermission("products:read"),
		handlers.Products,
	)

	// api.Get(
	// 	"/user/segma",
	// 	middleware.RequirePermission("products:read"),
	// 	handlers.Products,
	// )

	log.Fatal(app.Listen(":" + cfg.Port))
}
