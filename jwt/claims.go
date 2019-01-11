package jwt

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

// MakeClaims makes Claims structure to be used in API controller
func MakeClaims() jwt.Claims {
	return jwt.MapClaims{
		"foo": "bar",
		"iat": time.Now().UTC().Unix(),
	}
}
