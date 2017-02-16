package main

import (
	"flag"
	"fmt"

	"github.com/ikeikeikeike/cheapcdn/config"
	"github.com/ikeikeikeike/cheapcdn/gateway"
	"github.com/ikeikeikeike/cheapcdn/lib"
	"github.com/ikeikeikeike/cheapcdn/minio"

	"github.com/facebookgo/grace/gracehttp"
)

var cfg = &config.Config{
	Src:         *flag.String("src", "http://127.0.0.1:8888", "URL for own host"),
	Dest:        *flag.String("dest", "http://127.0.0.1:9000", "URL for proxy server"),
	AESSalt:     *flag.String("salt", "openunk-default-ses-saltown;pike", "ses salt"),
	GatewayUser: *flag.String("user", "user", "Set auth's username for issues apikey"),
	GatewayPass: *flag.String("pass", "pass", "Set auth's password for issues apikey"),
}

func main() {
	flag.Parse()

	lib.Load(cfg)
	gateway.Load(cfg)
	minio.Load(cfg)

	e := routes()
	lib.Routes(e)
	gateway.Routes(e)
	minio.Routes(e)

	e.Server.Addr = fmt.Sprintf(cfg.SrcHost())
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
