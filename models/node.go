//go:generate kallax gen
package models

import kallax "gopkg.in/src-d/go-kallax.v1"

// Node defines backend nodes.
type Node struct {
	kallax.Model `table:"nodes"`
	kallax.Timestamps

	ID       int64     `pk:"autoincr"`
	Objects  []*Object `fk:"node_id"`
	Host     string    `kallax:"host"`
	Free     int64     `kallax:"free"`
	Alive    bool      `kallax:"alive"`
	Provider string    `kallax:"provider"`
}
