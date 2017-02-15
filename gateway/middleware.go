package gateway

import "github.com/labstack/echo"

func basicAuth(user, pass string, c echo.Context) bool {
	if user == "joe" && pass == "secret" {
		return true
	}

	return false
}
