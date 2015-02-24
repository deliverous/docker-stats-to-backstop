// vim: ts=2:nowrap
package backstop

import (
	"net/http"
	"strings"
)

func SendMetrics(url string, metrics []Metric) (status int, err error) {
	data, err := RenderMetrics(metrics)
	if err != nil {
		return 0, err
	}

	client := http.Client{}
	response, err := client.Post(url, "application/json", strings.NewReader(data))
	if err != nil {
		return 0, err
	}
	status = response.StatusCode
	return
}
