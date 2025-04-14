package utils

import (
	"cisdi-technical-assessment/REST/auth-service/model"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTUtil interface {
	GenerateToken(user *model.User) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	ExtractUserID(token *jwt.Token) (uint, error)
}

type jwtUtil struct {
	secretKey string
	expiresIn time.Duration
}

func NewJWTUtil(secretKey string, expiresIn time.Duration) JWTUtil {
	return &jwtUtil{
		secretKey: secretKey,
		expiresIn: expiresIn,
	}
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func (j *jwtUtil) GenerateToken(user *model.User) (string, error) {
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtUtil) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtUtil) ExtractUserID(token *jwt.Token) (uint, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token claims")
	}

	userId, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id claim not found or invalid")
	}

	return uint(userId), nil
}
