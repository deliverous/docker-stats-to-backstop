// vim: ts=2 nowrap
package docker

import (
	"strings"
)

type Container struct {
	Id    string   ``
	Image string   ``
	Names []string ``
}

func (container Container) Name() string {
	var name string
	if len(container.Names) > 0 {
		name = container.Names[0]
	} else {
		name = container.Id
	}
	count := strings.Count(name, "/")

	for _, value := range container.Names {
		current_count := strings.Count(value, "/")
		if current_count < count {
			count = current_count
			name = value
		}
	}
	return name
}
