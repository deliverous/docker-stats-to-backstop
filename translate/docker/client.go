// vim: ts=2 nowrap
package docker

import (
	"encoding/json"
	"net/http"
)

func GetDockerStats(client *http.Client, urlStr string) (*ContainerStats, error) {
	request, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	stats := &ContainerStats{}
	err = decoder.Decode(stats)
	if err != nil {
		return nil, err
	}
	return stats, nil
}
