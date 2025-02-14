package controllers

import (
	"github.com/idir-44/mt5-cdn-project/internal/middlewares"
	"github.com/idir-44/mt5-cdn-project/internal/services"
	"github.com/idir-44/mt5-cdn-project/pkg/server"
)

type controller struct {
	service services.Service
}

func RegisterHandlers(routerGroup *server.Router, srv services.Service) {
	c := controller{srv}

	routerGroup.Use(middlewares.AddCurentUser)

	routerGroup.POST("/users", c.createUser)
	routerGroup.POST("/login", c.login)

	routerGroup.POST("/upload", c.uploadFile, middlewares.IsAuthenticated)
	routerGroup.GET("/download", c.downloadFile, middlewares.IsAuthenticated)

	routerGroup.GET("/me", c.getCurrentUser, middlewares.IsAuthenticated)

	// ✅ Ajout des routes manquantes
	routerGroup.GET("/files", c.listFiles, middlewares.IsAuthenticated)                // Liste les fichiers d'un dossier
	routerGroup.GET("/download-folder", c.downloadFolder, middlewares.IsAuthenticated) // Télécharge un dossier en ZIP
}
