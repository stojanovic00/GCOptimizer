package domain

import (
	"auth_service/errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type JwtClaims struct {
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
	jwt.StandardClaims
}

func (claims JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) {
		return nil
	}
	return errors.ErrInvalidToken{}
}
