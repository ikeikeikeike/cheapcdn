package gateway

import (
	"github.com/labstack/echo"
)

func basicAuth(user, pass string, ctx echo.Context) (error, bool) {
	if user == cfg.GatewayUser && pass == cfg.GatewayPass {
		return nil, true
	}

	return nil, false
}
