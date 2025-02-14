package controllers

import (
	"github.com/idir-44/mt5-cdn-project/internal/middlewares"
	"github.com/idir-44/mt5-cdn-project/internal/services"
	"github.com/idir-44/mt5-cdn-project/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type controller struct {
	service services.Service
}

func RegisterHandlers(routerGroup *server.Router, srv services.Service) {
	c := controller{srv}

	routerGroup.Use(middlewares.AddCurentUser)
	routerGroup.Use(middlewares.MetricsMiddleware)

	routerGroup.POST("/users", c.createUser)
	routerGroup.POST("/login", c.login)

	routerGroup.POST("/upload", c.uploadFile, middlewares.IsAuthenticated)
	routerGroup.GET("/download", c.downloadFile, middlewares.IsAuthenticated)
	routerGroup.GET("/me", c.getCurrentUser, middlewares.IsAuthenticated)

	routerGroup.GET("/files", c.listFiles, middlewares.IsAuthenticated)
	routerGroup.GET("/download-folder", c.downloadFolder, middlewares.IsAuthenticated)
	routerGroup.DELETE("/delete", c.deleteFileOrFolder, middlewares.IsAuthenticated)

	routerGroup.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}
