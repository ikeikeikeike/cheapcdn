package minio

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// NewMinoAdminReverseProxy returns as struct that direction for admin proxy.
func NewMinoAdminReverseProxy() *httputil.ReverseProxy {
	director := func(r *http.Request) {
		r.URL.Scheme = cfg.DestScheme()
		r.URL.Host = fmt.Sprintf(cfg.DestHost())
	}

	return &httputil.ReverseProxy{Director: director}
}

// NewMinoBucketReverseProxy returns as struct that direction for bucket proxy.
func NewMinoBucketReverseProxy() *httputil.ReverseProxy {
	director := func(r *http.Request) {
		r.URL.Scheme = cfg.DestScheme()
		r.URL.Host = fmt.Sprintf(cfg.DestHost())
	}

	return &httputil.ReverseProxy{Director: director}
}
