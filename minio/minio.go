package minio

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/ikeikeikeike/cheapcdn/lib"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewMinoAdminReverseProxy() *httputil.ReverseProxy {
	director := func(r *http.Request) {
		r.URL.Scheme = cfg.DestScheme()
		r.URL.Host = fmt.Sprintf(cfg.DestHost())
	}

	return &httputil.ReverseProxy{Director: director}
}

func NewMinoBucketReverseProxy(ctx echo.Context) *httputil.ReverseProxy {
	cc := ctx.(*lib.CacheContext)

	parts := strings.Split(cfg.KeyLookup, ":")
	extractor := keyFromHeader(parts[1], cfg.AuthScheme)
	switch parts[0] {
	case "query":
		extractor = keyFromQuery(parts[1])
	}

	director := func(r *http.Request) {
		// if config.Skipper(c) {
			// return next(c)
		// }

		// Extract and verify key
		key, err := extractor(ctx)
		if err != nil {
			log.Println("error=", err)
			return
		}
		if !validator(ctx, key) {
			return
		}

		r.URL.Scheme = cfg.DestScheme()
		r.URL.Host = fmt.Sprintf(cfg.DestHost())
	}

	return &httputil.ReverseProxy{Director: director}
}

func validator(ctx echo.Context, key) bool {
	return false
}

func keyFromHeader(header string, authScheme string) keyExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		if auth == "" {
			return "", errors.New("Missing key in request header")
		}
		if header == echo.HeaderAuthorization {
			l := len(authScheme)
			if len(auth) > l+1 && auth[:l] == authScheme {
				return auth[l+1:], nil
			}
			return "", errors.New("Invalid key in the request header")
		}
		return auth, nil
	}
}

func keyFromQuery(param string) keyExtractor {
	return func(c echo.Context) (string, error) {
		key := c.QueryParam(param)
		if key == "" {
			return "", errors.New("Missing key in the query string")
		}
		return key, nil
	}
}
