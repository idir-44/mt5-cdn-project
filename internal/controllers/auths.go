package controllers

import (
	"net/http"
	"time"

	"github.com/idir-44/mt5-cdn-project/internal/models"
	"github.com/labstack/echo/v4"
)

func (r controller) login(c echo.Context) error {
	req := models.LoginRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	user, token, err := r.service.Login(req)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"id":    user.ID,
		"email": user.Email,
		"token": token,
	})
}
