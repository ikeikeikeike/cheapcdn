package main

import (
	"strings"

	"github.com/labstack/echo"
	md "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func routes() *echo.Echo {
	e := echo.New()

	e.Logger.SetLevel(log.WARN)

	e.Use(md.LoggerWithConfig(md.LoggerConfig{
		Skipper: func(ctx echo.Context) bool {
			h := ctx.Request().Host
			if strings.HasPrefix(h, "localhost") {
				return false
			} else if strings.HasPrefix(h, "127.0.0.1") {
				return false
			}
			return true
		},
	}))

	e.Use(md.Recover())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "ok")
	})

	return e
}
