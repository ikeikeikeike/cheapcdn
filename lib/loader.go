package lib

import (
	"crypto/aes"
	"fmt"
	"time"

	"github.com/ikeikeikeike/cheapcdn/config"
	"github.com/labstack/echo"
	cache "github.com/patrickmn/go-cache"
)

var cfg *config.Config

// Load configuration instead of context.
func Load(c *config.Config) {
	var err error

	aesBlock, err = aes.NewCipher([]byte(c.AESSalt))
	if err != nil {
		panic(fmt.Sprintf("Error: NewCipher(%d bytes) = %s", len(c.AESSalt), err))
	}

	cfg = c
}

// Routes set handler into mux
func Routes(e *echo.Echo) {
	store := cache.New(5*time.Minute, 30*time.Second)

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			return h(&CacheContext{ctx, store})
		}
	})
}
