package minio

import (
	"github.com/ikeikeikeike/cheapcdn/config"
	"github.com/labstack/echo"
)

type (
	gateway struct {
		File   string `json:"f"`
		IPAddr string `json:"i"`
		Time   string `json:"t"`
		Node   string `json:"n"`
	}
)

const (
	authScheme = "Bearer"
	authParam  = "cdnkey"
)

var cfg *config.Config

// Load configuration instead of context.
func Load(c *config.Config) {
	cfg = c
}

// Routes set handler into mux
func Routes(e *echo.Echo) {
	g := e.Group("/")
	g.Use(keyAuth())
	g.Any("*", echo.WrapHandler(NewMinoBucketReverseProxy()))
}
