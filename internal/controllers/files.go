package controllers

import _ "fmt"

import (
	"github.com/idir-44/mt5-cdn-project/internal/middlewares"
	"net/http"
	"os"
	"path/filepath"

	"github.com/idir-44/mt5-cdn-project/internal/models"
	"github.com/labstack/echo/v4"
)

// Upload de fichier
func (r controller) uploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Fichier invalide"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Impossible d'ouvrir le fichier"})
	}
	defer src.Close()

	uploadDir := "uploads"
	os.MkdirAll(uploadDir, os.ModePerm)
	filePath := filepath.Join(uploadDir, file.Filename)

	dst, err := os.Create(filePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Erreur d'écriture du fichier"})
	}
	defer dst.Close()

	_, err = dst.ReadFrom(src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Erreur lors de la sauvegarde du fichier"})
	}

	user, err := middlewares.GetUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Utilisateur non authentifié"})
	}

	fileRecord := models.File{
		UserID:   user.ID,
		Filename: file.Filename,
		Filepath: filePath,
		Filesize: file.Size,
	}

	savedFile, err := r.service.UploadFile(fileRecord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Erreur d'enregistrement"})
	}

	return c.JSON(http.StatusOK, savedFile)
}

// Téléchargement de fichier
func (r controller) downloadFile(c echo.Context) error {
	fileID := c.QueryParam("id")

	file, err := r.service.GetFileByID(fileID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Fichier introuvable"})
	}

	return c.File(file.Filepath)
}
