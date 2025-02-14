package services

import (
	"github.com/idir-44/mt5-cdn-project/internal/models"
)

func (s service) UploadFile(req models.File) (models.File, error) {
	return s.repository.UploadFile(req)
}

func (s service) GetFileByID(id string) (models.File, error) {
	return s.repository.GetFileByID(id)
}

func (s service) DeleteFile(id string) error {
	return s.repository.DeleteFile(id)
}

func (s service) ListFiles(folderPath string) ([]models.File, error) {
	return s.repository.ListFiles(folderPath)
}
