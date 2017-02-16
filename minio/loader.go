package minio

import (
	"github.com/ikeikeikeike/cheapcdn/config"
	"github.com/labstack/echo"
)

var cfg *config.Config

// Load configuration instead of context.
func Load(c *config.Config) {
	cfg = c
}

// Routes set handler into mux
func Routes(e *echo.Echo) {
	ctx := e.AcquireContext()

	// For admin
	g := e.Group("/minio")
	g.Any("*", echo.WrapHandler(NewMinoAdminReverseProxy()))

	// For bucket
	g = e.Group("/")
	// g.Use(keyAuth)
	g.Any("*", echo.WrapHandler(NewMinoBucketReverseProxy(ctx)))
}
