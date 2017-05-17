package config

import (
	"database/sql"
	"net/url"
)

// Config defines for configuration
type Config struct {
	Src         string
	AESSalt     string
	GatewayUser string
	GatewayPass string
	DB          *sql.DB
}

// SrcHost returns URL host for this proxy.
func (cfg *Config) SrcHost() string {
	u, _ := url.Parse(cfg.Src)
	return u.Host
}
