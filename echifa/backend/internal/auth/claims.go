package auth

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	RealmAccess RealmAccess `json:"realm_access"`

	ResourceAccess map[string]ClientAccess `json:"resource_access"`

	jwt.RegisteredClaims
}

type RealmAccess struct {
	Roles []string `json:"roles"`
}

type ClientAccess struct {
	Roles []string `json:"roles"`
}

func (c *Claims) HasRealmRole(role string) bool {
	for _, r := range c.RealmAccess.Roles {
		if r == role {
			return true
		}
	}

	return false
}
