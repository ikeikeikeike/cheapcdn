package gateway

import (
	"github.com/labstack/echo"
)

func basicAuth(user, pass string, ctx echo.Context) bool {
	if user == cfg.GatewayUser && pass == cfg.GatewayPass {
		return true
	}

	return false
}
