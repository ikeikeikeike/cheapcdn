package minio

import (
	"github.com/labstack/echo"
	md "github.com/labstack/echo/middleware"
)

var keyAuth = md.KeyAuthWithConfig(md.KeyAuthConfig{
	KeyLookup: "query:authkey",
	Validator: func(key string, c echo.Context) bool {
		return key == "valid-key"
	},
})
