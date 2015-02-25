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
	version := struct{ ApiVersion string }{}
	if err := api.get(api.BaseUrl+"/version", &version); err != nil {
		return "", err
	}
	return "v" + version.ApiVersion, nil
}

func (api *DockerApi) GetContainers() ([]Container, error) {
	containers := []Container{}
	if err := api.get(api.url("containers", "json"), &containers); err != nil {
		return nil, err
	}
	return containers, nil
}

func (api *DockerApi) GetContainerStats(container string) (*ContainerStats, error) {
	stats := &ContainerStats{}
	if err := api.get(api.url("containers", container, "stats"), stats); err != nil {
		return nil, err
	}
	return stats, nil
}

func (api *DockerApi) get(urlStr string, data interface{}) error {
	request, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}

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
