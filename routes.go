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
		Skipper: func(ctx echo.Context) bool {
			if strings.HasPrefix(ctx.Request().Host, "localhost") {
				return true
			}
			return false
		},
	}))

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "ok")
	})

	return e
}
