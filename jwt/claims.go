package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// MakeClaims makes Claims structure to be used in API controller
func MakeClaims() jwt.Claims {
	return jwt.MapClaims{
		"foo": "bar",
		"iat": time.Now().UTC().Unix(),
	}
}
