package services

import (
	"github.com/idir-44/mt5-cdn-project/internal/models"
	"github.com/idir-44/mt5-cdn-project/internal/repositories"
)

type service struct {
	repository repositories.Repository
}

func NewService(repo repositories.Repository) Service {
	return service{repo}
}

type Service interface {
	CreateUser(req models.CreateUserReqesut) (models.User, error)
	Login(req models.LoginRequest) (models.User, string, error)

	UploadFile(req models.File) (models.File, error)
	GetFileByID(id string) (models.File, error)
	DeleteFile(id string) error
	ListFiles(folderPath string) ([]models.File, error)
}
