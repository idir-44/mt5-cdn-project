package controllers

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/idir-44/mt5-cdn-project/internal/middlewares"
	"github.com/idir-44/mt5-cdn-project/internal/models"
	"github.com/labstack/echo/v4"
)

func (r controller) uploadFile(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Requête invalide"})
	}

	files := form.File["files"]
	folderPath := c.FormValue("folder")

	if len(folderPath) > 255 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Chemin du dossier trop long"})
	}

	uploadDir := filepath.Join("uploads", folderPath)
	os.MkdirAll(uploadDir, os.ModePerm)

	user, err := middlewares.GetUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Utilisateur non authentifié"})
	}

	var savedFiles []models.File
	for _, file := range files {
		filePath := filepath.Join(uploadDir, file.Filename)

		src, err := file.Open()
		if err != nil {
			continue
		}
		defer src.Close()

		dst, err := os.Create(filePath)
		if err != nil {
			continue
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			continue
		}

		fileRecord := models.File{
			UserID:     user.ID,
			Filename:   file.Filename,
			FolderPath: folderPath,
			Filepath:   filePath,
			Filesize:   file.Size,
		}

		savedFile, err := r.service.UploadFile(fileRecord)
		if err == nil {
			savedFiles = append(savedFiles, savedFile)
		}
	}

	return c.JSON(http.StatusOK, savedFiles)
}

func (r controller) listFiles(c echo.Context) error {
	folderPath := c.QueryParam("folder")
	files, err := r.service.ListFiles(folderPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Erreur lors de la récupération des fichiers"})
	}
	return c.JSON(http.StatusOK, files)
}

func (r controller) downloadFile(c echo.Context) error {
	fileID := c.QueryParam("id")

	file, err := r.service.GetFileByID(fileID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Fichier introuvable"})
	}

	return c.File(file.Filepath)
}

func (r controller) downloadFolder(c echo.Context) error {
	folderPath := c.QueryParam("folder")
	zipPath := filepath.Join("uploads", folderPath+".zip")

	err := zipFolder(filepath.Join("uploads", folderPath), zipPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Erreur lors de la compression du dossier"})
	}

	return c.File(zipPath)
}

func (r controller) deleteFileOrFolder(c echo.Context) error {
	path := c.QueryParam("path")
	absolutePath := filepath.Join("uploads", path)

	info, err := os.Stat(absolutePath)
	if os.IsNotExist(err) {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Fichier ou dossier introuvable"})
	}

	if info.IsDir() {
		err = os.RemoveAll(absolutePath)
	} else {
		err = os.Remove(absolutePath)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Erreur lors de la suppression"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Suppression réussie"})
}

func zipFolder(source, target string) error {
	zipFile, err := os.Create(target)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du fichier ZIP: %w", err)
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Erreur lors du parcours du fichier:", err)
			return err
		}

		relativePath, _ := filepath.Rel(source, path)
		depth := strings.Count(relativePath, string(os.PathSeparator))
		if depth > 10 {
			fmt.Println("⚠️ Ignoré: trop profond:", relativePath)
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			fmt.Println("❌ Erreur création header ZIP:", err)
			return err
		}
		header.Name = relativePath

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			fmt.Println("❌ Erreur écriture ZIP:", err)
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				fmt.Println("❌ Impossible d'ouvrir le fichier:", path)
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				fmt.Println("❌ Erreur lors de la copie du fichier dans le ZIP:", path, err)
				return err
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("❌ Erreur lors de la création du ZIP:", err)
	}
	return err
}
