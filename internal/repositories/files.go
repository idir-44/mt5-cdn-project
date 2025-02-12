package repositories

import (
	"context"
	"time"

	"github.com/idir-44/mt5-cdn-project/internal/models"
)

func (r repository) UploadFile(file models.File) (models.File, error) {
	file.CreatedAt = time.Now().UTC()
	file.UpdatedAt = time.Now().UTC()

	_, err := r.db.NewInsert().Model(&file).ExcludeColumn("id").Returning("*").Exec(context.TODO())
	if err != nil {
		return models.File{}, err
	}
	return file, nil
}

func (r repository) GetFileByID(id string) (models.File, error) {
	file := models.File{}
	err := r.db.NewSelect().Model(&file).Where("id = ?", id).Scan(context.TODO())
	return file, err
}

func (r repository) DeleteFile(id string) error {
	_, err := r.db.NewDelete().Model((*models.File)(nil)).Where("id = ?", id).Exec(context.TODO())
	return err
}
