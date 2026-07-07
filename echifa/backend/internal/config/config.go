package config

import "os"

type Config struct {
	Port           string
	KeycloakIssuer string
	KeycloakJWKS   string
}

func Load() Config {
	return Config{
		Port:           getEnv("PORT", "8900"),
		KeycloakIssuer: getEnv("KEYCLOAK_ISSUER", "http://167.86.79.16:8080/realms/cnas-sso"),
		KeycloakJWKS:   getEnv("KEYCLOAK_JWKS_URL", "http://167.86.79.16:8080/realms/cnas-sso/protocol/openid-connect/certs"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
