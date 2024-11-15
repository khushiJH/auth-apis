package utils

import (
	"auth-api/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &models.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

func ParseToken(tokenString string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
