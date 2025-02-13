package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/idir-44/mt5-cdn-project/internal/controllers"
	"github.com/idir-44/mt5-cdn-project/internal/middlewares"
	"github.com/idir-44/mt5-cdn-project/internal/repositories"
	"github.com/idir-44/mt5-cdn-project/internal/services"
	"github.com/idir-44/mt5-cdn-project/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	e := echo.New()

	db, err := database.Connect()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewRepository(db)
	service := services.NewService(repo)

	v1 := e.Group("/v1")
	v1.Use(middlewares.AddCurentUser)
	controllers.RegisterHandlers(v1, service)

	req := httptest.NewRequest(http.MethodPost, "/v1/login", bytes.NewBufferString(`{"email": "test@test.com", "password": "password"}`))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
