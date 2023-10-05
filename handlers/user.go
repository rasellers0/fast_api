package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Fuck you!")
}

func GetUserData(c echo.Context) error {
	usr := &User

	return c.String(http.StatusOK, "No User Data")
}
