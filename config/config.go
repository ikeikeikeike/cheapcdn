package config

import "net/url"

// Config defines for configuration
type Config struct {
	Nodes       []string  // Launched cheapcdn servers
	Src         string
	Dest        string
	AESSalt     string
	GatewayUser string
	GatewayPass string
}

// SrcHost returns URL host for this proxy.
func (cfg *Config) SrcHost() string {
	u, _ := url.Parse(cfg.Src)
	return u.Host
}

// DestHost returns URL host for destnation system.
func (cfg *Config) DestHost() string {
	u, _ := url.Parse(cfg.Dest)
	return u.Host
}

// DestScheme returns URL scheme for destnation system.
func (cfg *Config) DestScheme() string {
	u, _ := url.Parse(cfg.Dest)
	return u.Scheme
}
