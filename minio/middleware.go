package minio

import (
	"errors"
	"net/http"

	"github.com/ikeikeikeike/cheapcdn/lib"
	"github.com/labstack/echo"
)

const (
	authScheme = "Bearer"
	parameter  = "cheapkey"
)

func keyAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Extract and verify key
			key, err := extractor(ctx)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden)
			}
			if !validator(ctx, key) {
				return echo.ErrUnauthorized
			}

			return next(ctx)
		}
	}
}

func validator(ctx echo.Context, key string) bool {
	_, b := ctx.(*lib.CacheContext).Store.Get(key)
	if !b {
		return false
	}

	// ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	// jsonstr.

	return true
}

func extractor(ctx echo.Context) (string, error) {
	token, err := extractHeader(ctx)
	if err == nil {
		return token, nil
	}

	token, err = extractParam(ctx)
	if err == nil {
		return token, nil
	}

	return "", errors.New("Missing token")
}

func extractHeader(ctx echo.Context) (string, error) {
	auth := ctx.Request().Header.Get(echo.HeaderAuthorization)
	if auth == "" {
		return "", errors.New("Missing token in request header")
	}

	length := len(authScheme)

	if len(auth) > length+1 && auth[:length] == authScheme {
		ctx.Request().Header.Del(echo.HeaderAuthorization)
		return auth[length+1:], nil
	}

	return "", errors.New("Invalid token in the request header")
}

func extractParam(ctx echo.Context) (string, error) {
	key := ctx.QueryParam(parameter)
	if key == "" {
		return "", errors.New("Missing token in the query string")
	}

	return key, nil
}
