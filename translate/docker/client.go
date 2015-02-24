// vim: ts=2 nowrap
package docker

import (
	"gopkg.in/jmcvetta/napping.v1"
)

func GetDockerStats(session *napping.Session, url string, stats *ContainerStats) (int, error) {
	response, err := session.Get(url, nil, stats, nil)
	if err != nil {
		return 0, err
	}
	return response.Status(), nil
}
