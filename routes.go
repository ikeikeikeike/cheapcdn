package main

import (
	"strings"

	"github.com/labstack/echo"
	md "github.com/labstack/echo/middleware"
)

func routes() *echo.Echo {
	e := echo.New()
	e.Use(md.Logger())
	e.Use(md.Recover())
	e.Use(md.LoggerWithConfig(md.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Request().Host, "localhost") {
				return true
			}
			return false
		},
	}))

	return e
}
