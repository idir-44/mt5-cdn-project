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
	// ✅ Initialisation de la base de données en mémoire pour le test
	db, err := database.Connect()
	if err != nil {
		t.Fatalf("failed to connect to database: %s", err)
	}
	defer db.Close()

	repo := repositories.NewRepository(db)

	file := models.File{
		ID:        "1234",
		UserID:    "5678",
		Filename:  "testfile.txt",
		Filepath:  "/files/testfile.txt",
		Filesize:  1024,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// ✅ Test de l'upload d'un fichier
	uploadedFile, err := repo.UploadFile(file)
	assert.NoError(t, err, "L'upload du fichier a échoué")
	assert.Equal(t, file.Filename, uploadedFile.Filename, "Les noms de fichiers ne correspondent pas")
}
