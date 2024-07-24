package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int  `json:"user_id"`
	Role   int8 `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, secretKey string, role int8) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func GetJWTClaims(tokenString, secretKey string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
