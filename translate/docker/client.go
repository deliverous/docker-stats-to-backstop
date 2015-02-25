// vim: ts=2 nowrap
package docker

import (
	"encoding/json"
	"net/http"
	"path"
)

const (
	ApiVersion = "v1.17"
)

type DockerApi struct {
	Client  *http.Client
	BaseUrl string
}

func NewDockerApi(client *http.Client, baseUrl string) *DockerApi {
	return &DockerApi{
		Client:  client,
		BaseUrl: baseUrl,
	}
}

func (api *DockerApi) GetApiVersion() (string, error) {
	request, err := http.NewRequest("GET", api.BaseUrl+"/version", nil)
	if err != nil {
		return "", err
	}

	version := struct{ ApiVersion string }{}
	if err := api.processRequest(request, &version); err != nil {
		return "", err
	}
	return "v" + version.ApiVersion, nil
}

func (api *DockerApi) GetContainers() ([]Container, error) {
	request, err := http.NewRequest("GET", api.url("containers", "json"), nil)
	if err != nil {
		return nil, err
	}

	containers := []Container{}
	if err := api.processRequest(request, &containers); err != nil {
		return nil, err
	}
	return containers, nil
}

func (api *DockerApi) GetContainerStats(container string) (*ContainerStats, error) {
	request, err := http.NewRequest("GET", api.url("containers", container, "stats"), nil)
	if err != nil {
		return nil, err
	}

	stats := &ContainerStats{}
	if err := api.processRequest(request, stats); err != nil {
		return nil, err
	}
	return stats, nil
}

func (api *DockerApi) processRequest(request *http.Request, data interface{}) error {
	response, err := api.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode < 300 {
		return json.NewDecoder(response.Body).Decode(data)
	}
	return nil
}

func (api *DockerApi) url(parts ...string) string {
	return path.Join(api.BaseUrl, ApiVersion, path.Join(parts...))
}
