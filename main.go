package main

import (
	"flag"
	"fmt"

	"github.com/ikeikeikeike/authproxy/authproxy"
	"github.com/ikeikeikeike/authproxy/gateway"
	"github.com/ikeikeikeike/authproxy/minio"

	"github.com/facebookgo/grace/gracehttp"
)

var port = flag.Int("port", 8888, "port number")
var dest = flag.Int("dest", 9000, "port number for proxy server")
var salt = flag.String("s", "openunk-default-ses-saltown;pike", "ses salt")

func main() {
	flag.Parse()
	minio.Load(*dest)
	authproxy.Load(*salt)

	e := routes()
	gateway.Routes(e)
	minio.Routes(e)

	e.Server.Addr = fmt.Sprintf(":%v", *port)
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
