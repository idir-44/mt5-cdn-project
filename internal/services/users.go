package services

import (
	"github.com/idir-44/mt5-cdn-project/internal/models"
	"github.com/idir-44/mt5-cdn-project/pkg/utils"
)

func (s service) CreateUser(req models.CreateUserReqesut) (models.User, error) {

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return models.User{}, err
	}

	user, err := s.repository.CreateUser(models.User{Email: req.Email, Password: password})

	return user, err
}
