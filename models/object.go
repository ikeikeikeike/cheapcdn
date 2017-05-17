//go:generate kallax gen
package models

import kallax "gopkg.in/src-d/go-kallax.v1"

// Object defines that file ops.
type Object struct {
	kallax.Model `table:"objects"`
	kallax.Timestamps

	ID   int64  `pk:"autoincr"`
	Node *Node  `fk:"node_id,inverse"`
	Name string `kallax:"name"`
}
