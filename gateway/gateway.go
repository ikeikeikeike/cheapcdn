package gateway

import (
	"net/http"

	"github.com/labstack/echo"
	md "github.com/labstack/echo/middleware"
)

func gateway(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}

func Routes(e *echo.Echo) {
	g := e.Group("/gateway")
	g.Use(md.BasicAuth(basicAuth))
	g.GET("/", gateway)
}
