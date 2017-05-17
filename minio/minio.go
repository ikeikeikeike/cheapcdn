package minio

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"

	"github.com/ikeikeikeike/cheapcdn/lib"
)

// NewMinoBucketReverseProxy returns as struct that direction for bucket proxy.
func NewMinoBucketReverseProxy() *httputil.ReverseProxy {
	director := func(r *http.Request) {
		key := r.URL.Query().Get(authParam)

		var g *gateway
		json.Unmarshal(lib.DecryptAexHex(key), &g)

		r.URL.Scheme = "http"
		r.URL.Host = g.Node
	}

	return &httputil.ReverseProxy{Director: director}
}
