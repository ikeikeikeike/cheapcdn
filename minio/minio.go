package minio

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/labstack/echo"
)

var dest int

func director(r *http.Request) {
	r.URL.Scheme = "http"
	r.URL.Host = fmt.Sprintf(":%v", dest)
}

// Routes set handler into mux
func Routes(e *echo.Echo) {
	// For admin
	g := e.Group("/minio")
	g.Any("*", echo.WrapHandler(&httputil.ReverseProxy{Director: director}))

	// For bucket
	g = e.Group("/")
	g.Use(keyAuth)
	g.Any("*", echo.WrapHandler(&httputil.ReverseProxy{Director: director}))
}

// Load configuration instead of context.
func Load(port int) {
	dest = port
}
