// vim: ts=2:nowrap
package server

import (
	"github.com/deliverous/docker-stats-to-backstop/translate"
	"github.com/deliverous/docker-stats-to-backstop/translate/backstop"
	"github.com/deliverous/docker-stats-to-backstop/translate/docker"
	"log"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	Hostname    string
	DockerURL   string
	BackstopURL string
	Poll        time.Duration
	Rules       *Rules
	Verbose     bool
}

func (server *Server) Serve() {
	transport := &http.Transport{}
	transport.RegisterProtocol("unix", &SocketTransport{
		Predicate:         LstatSocketPredicate,
		Timeout:           32 * time.Second,
		DisableKeepAlives: true,
	})
	client := &http.Client{Transport: transport, Timeout: 32 * time.Second}

	dockerApi := docker.NewDockerApi(client, server.DockerURL, server.Verbose)

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
			categories := make(map[string]int64)

			metrics := []backstop.Metric{}
			log.Printf("Processing %d containers", len(containers))
			now := time.Now().Unix()

			for _, container := range containers {
				prefix, category := ApplyRules(server.Rules, container.Name())

				log.Printf("Processing container %s (%s: %s)\n", container.Id[:12], prefix, category)
				stats, err := dockerApi.GetContainerStats(container.Id)
				if err != nil {
					log.Printf("ERROR: cannot get container stats: %s\n", err)
					continue
				}
				metrics = append(metrics, translate.Translate(prefix, stats)...)
				categories[category] += 1
			}

			total := int64(0)
			for category, value := range categories {
				metrics = append(metrics, backstop.Metric{
					Name:      server.Hostname + ".containers." + category,
					Value:     value,
					Timestamp: now,
				})
				total += value
			}
			metrics = append(metrics, backstop.Metric{
				Name:      server.Hostname + ".containers.total",
				Value:     total,
				Timestamp: now,
			})

			err = backstop.SendMetrics(client, server.BackstopURL, metrics, server.Verbose)
			if err != nil {
				log.Printf("ERROR: cannot send stats: %s\n", err)
			}
		}

		time.Sleep(server.Poll)
	}

}

func reverseHostname(hostname string) string {
	parts := strings.Split(hostname, ".")
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	return strings.Join(parts, ".")
}
