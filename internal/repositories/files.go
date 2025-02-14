package repositories

import (
	"context"
	"time"

	"github.com/idir-44/mt5-cdn-project/internal/models"
)

// Upload du fichier
func (r repository) UploadFile(file models.File) (models.File, error) {
	file.CreatedAt = time.Now().UTC()
	file.UpdatedAt = time.Now().UTC()

	_, err := r.db.NewInsert().Model(&file).ExcludeColumn("id").Returning("*").Exec(context.TODO())
	if err != nil {
		return models.File{}, err
	}
	return file, nil
}

// Récupérer un fichier par son ID
func (r repository) GetFileByID(id string) (models.File, error) {
	file := models.File{}
	err := r.db.NewSelect().Model(&file).Where("id = ?", id).Scan(context.TODO())
	return file, err
}

// Supprimer un fichier par son ID
func (r repository) DeleteFile(id string) error {
	_, err := r.db.NewDelete().Model((*models.File)(nil)).Where("id = ?", id).Exec(context.TODO())
	return err
}

// ✅ Liste les fichiers d'un dossier
func (r repository) ListFiles(folder string) ([]models.File, error) {
	var files []models.File
	err := r.db.NewSelect().Model(&files).Where("folder_path = ?", folder).Scan(context.TODO())
	return files, err
}
