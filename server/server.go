// vim: ts=2:nowrap
package server

import (
	"github.com/deliverous/docker-stats-to-backstop/translate"
	"github.com/deliverous/docker-stats-to-backstop/translate/backstop"
	"github.com/deliverous/docker-stats-to-backstop/translate/docker"
	"log"
	"net/http"
	"time"
)

func ServeForever(dockerUrl string, backstopUrl string, prefix string, duration time.Duration) {
	prefixRule, err := loadPrefixRule(prefix)
	if err != nil {
		log.Fatalf("ERROR: cannot load prefix rules: %s", err)
	}

	transport := &http.Transport{}
	transport.RegisterProtocol("unix", NewSocketTransport(LstatSocketPredicate, 2*time.Second))
	client := &http.Client{Transport: transport}

	dockerApi := docker.NewDockerApi(client, dockerUrl)

	version, err := dockerApi.GetApiVersion()
	if err != nil {
		log.Fatalf("ERROR: cannot get docker API version: %s", err)
	}

	if version != docker.ApiVersion {
		log.Println("WARNING: not using the latest api version")
		log.Printf("Using  : '%s'\n", docker.ApiVersion)
		log.Printf("Current: '%s'\n", version)
	}

	for {
		if containers, err := dockerApi.GetContainers(); err != nil {
			log.Printf("ERROR: cannot get docker containers list: %s", err)
		} else {
			for _, container := range containers {
				prefix := computePrefix(&container, prefixRule)

				log.Printf("Processing container %s (%s)\n", container.Id[:12], prefix)
				stats, err := dockerApi.GetContainerStats(container.Id)
				if err != nil {
					log.Printf("ERROR: cannot get container stats: %s\n", err)
					continue
				}
				err = backstop.SendMetrics(client, backstopUrl, translate.Translate(prefix, stats))
				if err != nil {
					log.Printf("ERROR: cannot send container stats: %s\n", err)
				}
			}
		}

		time.Sleep(duration)
	}
}
