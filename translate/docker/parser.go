// vim: ts=2 nowrap
package docker

import (
	"encoding/json"
)

func ParseDockerStats(data string) (*ContainerStats, error) {
	var content *ContainerStats
	err := json.Unmarshal([]byte(data), &content)
	return content, err
}
