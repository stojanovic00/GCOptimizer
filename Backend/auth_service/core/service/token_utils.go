package service

import (
	"auth_service/core/domain"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var JwtSecretKey = os.Getenv("JWT_SECRET_KEY")

func GenerateToken(account *domain.Account) (string, error) {
	//Custom claims
	claims := domain.JwtClaims{
		Email: account.Email,
		Role:  account.Role.Name,
	}

	//Standard claims
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.ExpiresAt = time.Now().Add(time.Hour).UTC().Unix()

	//Generates not yet signed token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*domain.JwtClaims, error) {
	claims := &domain.JwtClaims{}
	token, err := getTokenFromString(tokenString, claims)

	if err != nil {
		return claims, err
	}

	if token.Valid {
		if e := claims.Valid(); e == nil {
			return claims, e
		}
	}

	return claims, nil
}

func getTokenFromString(tokenString string, claims *domain.JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(JwtSecretKey), nil
		})
}
