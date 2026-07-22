package utils

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("campus-course-system-secret-key-2026")

type CustomClaims struct {
	UserID    uint64 `json:"user_id"`
	StudentNo string `json:"student_no"`
	Role      string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(
	userID uint64,
	studentNo string,
	role string,
) (string, error) {
	claims := CustomClaims{
		UserID:    userID,
		StudentNo: studentNo,
		Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "campus-course-server",
			Subject:   studentNo,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("无效的签名方法")
			}

			return jwtSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("无效的Token")
	}

	return claims, nil
}
