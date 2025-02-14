package repositories

import (
	"github.com/idir-44/mt5-cdn-project/internal/models"
	"github.com/idir-44/mt5-cdn-project/pkg/database"
)

type repository struct {
	db *database.DBConnection
}

func NewRepository(db *database.DBConnection) Repository {
	return repository{db}
}

type Repository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUser(id string) (models.User, error)

	UploadFile(file models.File) (models.File, error)
	GetFileByID(id string) (models.File, error)
	DeleteFile(id string) error
	ListFiles(folder string) ([]models.File, error) // ✅ Ajout
}
