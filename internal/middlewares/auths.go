package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/idir-44/mt5-cdn-project/internal/jwttoken"
	"github.com/idir-44/mt5-cdn-project/internal/models"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) (models.User, error) {

	user := c.Get("user")

	if user == nil {
		return models.User{}, fmt.Errorf("current user not found")
	}

	return user.(models.User), nil
}

func AddCurentUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookies, err := c.Cookie("token")
		if err != nil {
			return next(c)
		}

		key := os.Getenv("jwt_secret")
		if key == "" {
			return fmt.Errorf("jwt secret is not set")
		}
		user, err := jwttoken.ParseToken(cookies.Value, key)
		if err != nil {
			return err
		}

		c.Set("user", user)

		return next(c)

	}
}

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := GetUser(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		if user.ID == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		return next(c)
	}
}
