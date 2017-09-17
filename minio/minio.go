package minio

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/ikeikeikeike/cheapcdn/lib"
)

// NewMinoBucketReverseProxy returns as struct that direction for bucket proxy.
func NewMinoBucketReverseProxy() *httputil.ReverseProxy {
	director := func(r *http.Request) {
		key := r.URL.Query().Get(authParam)

		var g *gateway
		json.Unmarshal(lib.DecryptAexHex(key), &g)

		u, _ := url.Parse(g.Node)
		r.URL.Scheme = u.Scheme
		r.URL.Host = u.Host
		r.Host = u.Host
	}

	return &httputil.ReverseProxy{Director: director}
}
