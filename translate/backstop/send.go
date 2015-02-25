// vim: ts=2:nowrap
package backstop

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func SendMetrics(client *http.Client, urlStr string, metrics []Metric) error {
	payload, err := json.Marshal(&metrics)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	response.Body.Close()

	if response.StatusCode >= 400 {
		return errors.New(http.StatusText(response.StatusCode))
	}
	return nil
}
