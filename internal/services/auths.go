package services

import (
	"fmt"
	"os"

	"github.com/idir-44/mt5-cdn-project/internal/jwttoken"
	"github.com/idir-44/mt5-cdn-project/internal/models"
	"github.com/idir-44/mt5-cdn-project/pkg/utils"
)

func (s service) Login(req models.LoginRequest) (models.User, string, error) {
	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return models.User{}, "", err
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		return models.User{}, "", fmt.Errorf("invalid password: %s", err)
	}

	key := os.Getenv("jwt_secret")
	if key == "" {
		return models.User{}, "", fmt.Errorf("jwt secret is not set")
	}

	token, err := jwttoken.CreateToken(user, key, jwttoken.TokenTypeAccess)
	if err != nil {
		return models.User{}, "", err
	}

	return user, token, nil

}
