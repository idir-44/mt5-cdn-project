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

	// protected routes
	routerGroup.GET("/me", c.getCurrentUser, middlewares.IsAuthenticated)
}
