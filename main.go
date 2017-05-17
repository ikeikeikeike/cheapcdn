package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/lib/pq"

	einhorn "github.com/dcu/http-einhorn"
	"github.com/ikeikeikeike/cheapcdn/config"
	"github.com/ikeikeikeike/cheapcdn/gateway"
	"github.com/ikeikeikeike/cheapcdn/lib"
	"github.com/ikeikeikeike/cheapcdn/minio"
)

var (
	src  = flag.String("src", "http://127.0.0.1:8888", "URL for own host")
	salt = flag.String("salt", "openunk-default-ses-saltown;pike", "ses salt")
	dsn  = flag.String("dsn", "postgres://postgres:@127.0.0.1:5432/cheapcdn?sslmode=disable", "schema db uri")
	// dsn  = flag.String("dsn", "postgres:@tcp(127.0.0.1:3306)/cheapcdn?parseTime=true", "schema db uri")
	user = flag.String("user", "user", "Set auth's username for issues apikey")
	pass = flag.String("pass", "pass", "Set auth's password for issues apikey")
)

func main() {
	flag.Parse()

	db, err := sql.Open("postgres", *dsn)
	if err != nil {
		panic(fmt.Sprintf("It was unable to connect to the DB.\n%s\n", err))
	}

	cfg := &config.Config{
		Src:         *src,
		AESSalt:     *salt,
		GatewayUser: *user,
		GatewayPass: *pass,
		DB:          db,
	}

	lib.Load(cfg)
	gateway.Load(cfg)
	minio.Load(cfg)

	e := routes()
	lib.Routes(e)
	gateway.Routes(e)
	minio.Routes(e)

	e.Server.Addr = fmt.Sprintf(cfg.SrcHost())

	if einhorn.IsRunning() {
		einhorn.Run(e.Server, 0)
	} else {
		e.StartServer(e.Server)
	}
}
