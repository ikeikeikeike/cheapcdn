package main

import (
	"flag"
	"fmt"

	einhorn "github.com/dcu/http-einhorn"
	"github.com/ikeikeikeike/cheapcdn/config"
	"github.com/ikeikeikeike/cheapcdn/gateway"
	"github.com/ikeikeikeike/cheapcdn/lib"
	"github.com/ikeikeikeike/cheapcdn/minio"
	"github.com/k0kubun/pp"
)

var (
	src  = flag.String("src", "http://127.0.0.1:8888", "URL for own host")
	dest = flag.String("dest", "http://127.0.0.1:9000", "URL for proxy server")
	salt = flag.String("salt", "openunk-default-ses-saltown;pike", "ses salt")
	user = flag.String("user", "user", "Set auth's username for issues apikey")
	pass = flag.String("pass", "pass", "Set auth's password for issues apikey")
)

func main() {
	flag.Parse()
	cfg := &config.Config{
		Src:         *src,
		Dest:        *dest,
		AESSalt:     *salt,
		GatewayUser: *user,
		GatewayPass: *pass,
	}

	lib.Load(cfg)
	gateway.Load(cfg)
	minio.Load(cfg)

	e := routes()
	lib.Routes(e)
	gateway.Routes(e)
	minio.Routes(e)

	pp.Println(cfg)
	e.Server.Addr = fmt.Sprintf(cfg.SrcHost())

	if einhorn.IsRunning() {
		einhorn.Run(e.Server, 0)
	} else {
		e.StartServer(e.Server)
	}
}
