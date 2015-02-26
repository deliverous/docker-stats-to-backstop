// vim: ts=2:nowrap
package backstop

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/deliverous/docker-stats-to-backstop/utils"
	"log"
	"net/http"
)

func SendMetrics(client *http.Client, urlStr string, metrics []Metric, verbose bool) error {
	payload, err := json.Marshal(&metrics)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	if verbose {
		log.Println("--------------------------------------------------------------------------------")
		log.Println("REQUEST")
		log.Println("--------------------------------------------------------------------------------")
		log.Println(utils.PrettyPrint(request))
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	response.Body.Close()

	if verbose {
		log.Println("--------------------------------------------------------------------------------")
		log.Println("RESPONSE")
		log.Println("--------------------------------------------------------------------------------")
		log.Println(utils.PrettyPrint(response))
	}

	if response.StatusCode >= 400 {
		return errors.New(http.StatusText(response.StatusCode))
	}
	return nil
}
