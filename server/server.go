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

	var channel = make(chan byte, 20)
	for {
		go func() {
			channel <- 1
			server.loop(client, dockerApi, len(channel))
			<-channel
		}()

		time.Sleep(server.Poll)
	}

}

func (server *Server) loop(client *http.Client, dockerApi *docker.DockerApi, concurrency int) {
	start := time.Now()

	if containers, err := dockerApi.GetContainers(); err != nil {
		log.Printf("ERROR: cannot get docker containers list: %s", err)
	} else {
		categories := make(map[string]uint64)

		metrics := []backstop.Metric{}
		now := time.Now()

		for _, container := range containers {
			prefix, category := ApplyRules(server.Rules, container.Name())

			log.Printf("Processing container %s ('%s': '%s')\n", container.Id[:12], prefix, category)
			stats, err := dockerApi.GetContainerStats(container.Id)
			if err != nil {
				log.Printf("ERROR: cannot get container stats: %s\n", err)
				continue
			}
			metrics = append(metrics, translate.TranslateStats(prefix, stats)...)
			json, err := dockerApi.GetContainerJson(container.Id)
			if err != nil {
				log.Printf("ERROR: cannot get container json: %s\n", err)
				continue
			}
			metrics = append(metrics, translate.TranslateJson(prefix, json, now)...)
			if category != "" {
				categories[category] += 1
			}
		}

		delta := time.Now().Sub(start).Nanoseconds()

		log.Printf("Container total: %d\n", len(containers))

		for category, value := range categories {
			log.Printf("Container %s: %d\n", category, value)
			metrics = append(metrics, backstop.Metric{
				Name:      server.Hostname + ".containers." + category,
				Value:     value,
				Timestamp: now.Unix(),
			})
		}
		metrics = append(metrics, backstop.Metric{
			Name:      server.Hostname + ".containers.total",
			Value:     uint64(len(containers)),
			Timestamp: now.Unix(),
		})

		metrics = append(metrics, backstop.Metric{
			Name:      server.Hostname + ".metrics.threads",
			Value:     uint64(concurrency),
			Timestamp: now.Unix(),
		})

		metrics = append(metrics, backstop.Metric{
			Name:      server.Hostname + ".metrics.duration",
			Value:     uint64(delta),
			Timestamp: now.Unix(),
		})

		err = backstop.SendMetrics(client, server.BackstopURL, metrics, server.Verbose)
		if err != nil {
			log.Printf("ERROR: cannot send stats: %s\n", err)
		}
	}
}

func reverseHostname(hostname string) string {
	parts := strings.Split(hostname, ".")
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	return strings.Join(parts, ".")
}
