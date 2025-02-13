package tests

import (
	"testing"
	"time"

	"github.com/idir-44/mt5-cdn-project/internal/models"
	"github.com/idir-44/mt5-cdn-project/internal/repositories"
	"github.com/idir-44/mt5-cdn-project/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestUploadFile(t *testing.T) {
	db, err := database.Connect()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewRepository(db)

	file := models.File{
		UserID:    "5678",
		Filename:  "testfile.txt",
		Filepath:  "/files/testfile.txt",
		Filesize:  1024,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	uploadedFile, err := repo.UploadFile(file)
	assert.NoError(t, err)
	assert.Equal(t, file.Filename, uploadedFile.Filename)
}
