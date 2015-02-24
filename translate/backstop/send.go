// vim: ts=2:nowrap
package backstop

import (
	"gopkg.in/jmcvetta/napping.v1"
)

func SendMetrics(session *napping.Session, url string, metrics []Metric) (int, error) {
	response, err := session.Post(url, metrics, nil, nil)
	if err != nil {
		return 0, err
	}
	return response.Status(), nil
}
