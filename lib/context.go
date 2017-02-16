package lib

import (
	"github.com/labstack/echo"
	cache "github.com/patrickmn/go-cache"
)

type CacheContext struct {
	echo.Context
	Store *cache.Cache
}
