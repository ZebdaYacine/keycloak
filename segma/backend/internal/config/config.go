package config

import "os"

type Config struct {
	Port           string
	KeycloakIssuer string
	KeycloakJWKS   string
}

func Load() Config {
	return Config{
		Port:           getEnv("PORT", "8901"),
		KeycloakIssuer: getEnv("KEYCLOAK_ISSUER", "http://127.0.0.1:8080/realms/cnas-sso"),
		KeycloakJWKS:   getEnv("KEYCLOAK_JWKS_URL", "http://127.0.0.1:8080/realms/cnas-sso/protocol/openid-connect/certs"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
