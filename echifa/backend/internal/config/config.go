package config

import "os"

type Config struct {
	Port           string
	KeycloakIssuer string
	KeycloakJWKS   string
}

func Load() Config {
	return Config{
		Port:           mustGetEnv("PORT"),
		KeycloakIssuer: mustGetEnv("KEYCLOAK_ISSUER"),
		KeycloakJWKS:   mustGetEnv("KEYCLOAK_JWKS_URL"),
	}
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " environment variable is required")
	}
	return value
}
