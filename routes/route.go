package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is Product API")
	})
}
