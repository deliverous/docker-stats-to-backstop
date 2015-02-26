// vim: ts=2 nowrap
package docker

import (
	"encoding/json"
	"errors"
	"github.com/deliverous/docker-stats-to-backstop/utils"
	"log"
	"net/http"
	"path"
)

const (
	ApiVersion = "v1.17"
)

type DockerApi struct {
	Client  *http.Client
	BaseUrl string
	Verbose bool
}

func NewDockerApi(client *http.Client, baseUrl string, verbose bool) *DockerApi {
	return &DockerApi{
		Client:  client,
		BaseUrl: baseUrl,
		Verbose: verbose,
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

	if api.Verbose {
		log.Println("--------------------------------------------------------------------------------")
		log.Println("REQUEST")
		log.Println("--------------------------------------------------------------------------------")
		log.Println(utils.PrettyPrint(request))
	}

	response, err := api.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if api.Verbose {
		log.Println("--------------------------------------------------------------------------------")
		log.Println("RESPONSE")
		log.Println("--------------------------------------------------------------------------------")
		log.Println(utils.PrettyPrint(response))
	}

	if response.StatusCode >= 400 {
		return errors.New(http.StatusText(response.StatusCode))
	}
	if response.StatusCode < 300 {
		return json.NewDecoder(response.Body).Decode(data)
	}
	return nil
}

func (api *DockerApi) url(parts ...string) string {
	return api.BaseUrl + "/" + ApiVersion + "/" + path.Join(parts...)
}
