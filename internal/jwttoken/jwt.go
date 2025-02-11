package jwttoken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/idir-44/mt5-cdn-project/internal/models"
)

type TokenType string

const (
	TokenTypeAccess          TokenType = "accessToken"
	TokenTypeResetPassword   TokenType = "resetPassword"
	TokenTypeEmailValidation TokenType = "emailValidation"
)

type jwtClaims struct {
	models.User
	jwt.StandardClaims
}

func CreateToken(user models.User, key string, tokenType TokenType) (string, error) {

	var tokenExpiry time.Duration

	switch tokenType {
	case TokenTypeAccess:
		tokenExpiry = 24
	case TokenTypeResetPassword:
		tokenExpiry = 2
	case TokenTypeEmailValidation:
		tokenExpiry = 2
	default:
		tokenExpiry = 1
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * tokenExpiry).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        uuid.New().String(),
		},
	})

	return token.SignedString([]byte(key))
}

func ParseToken(tokenString, key string) (models.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		return claims.User, nil
	} else {
		return models.User{}, fmt.Errorf("error parsing token: %s", err)
	}
}
