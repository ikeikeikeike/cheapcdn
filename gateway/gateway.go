package gateway

import (
	"net/http"

	"github.com/labstack/echo"
)

func gateway(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Welcome!")
}
