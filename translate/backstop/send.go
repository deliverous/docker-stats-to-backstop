// vim: ts=2:nowrap
package backstop

import (
	"net/http"
	"strings"
)

func SendMetrics(url string, metrics []Metric, clientOptions []clientOption, requestOptions []requestOption) (status int, err error) {
	data, err := RenderMetrics(metrics)
	if err != nil {
		return 0, err
	}

	client := &http.Client{}
	for _, option := range clientOptions {
		option(client)
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return 0, err
	}
	for _, option := range requestOptions {
		option(request)
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	status = response.StatusCode
	return
}

type clientOption func(client *http.Client)
type requestOption func(request *http.Request)

func client(options ...clientOption) []clientOption {
	return options
}

func request(options ...requestOption) []requestOption {
	return options
}

func BasicAuth(username, password string) requestOption {
	return func(request *http.Request) {
		request.SetBasicAuth(username, password)
	}
}
