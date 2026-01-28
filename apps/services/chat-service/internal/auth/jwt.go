package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Sub uint `json:"sub"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string) (uint, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return 0, fmt.Errorf("JWT_SECRET chưa được cấu hình")
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims.Sub, nil
	}

	return 0, fmt.Errorf("token không hợp lệ")
}
