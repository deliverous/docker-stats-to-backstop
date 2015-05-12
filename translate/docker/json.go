// vim: ts=2 nowrap
package docker

import (
	"time"
)

type ContainerJson struct {
	State StateJson `json:"State"`
}

type StateJson struct {
	StartedAt time.Time `json:"StartedAt"`
}
