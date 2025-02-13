package controllers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
)

func ReverseProxy(c echo.Context) error {
	backendURL := os.Getenv("BACKEND_URL")
	if backendURL == "" {
		backendURL = "http://localhost:8081"
	}

	target, err := url.Parse(backendURL)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Invalid backend URL")
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(c.Response().Writer, c.Request())
	return nil
}
