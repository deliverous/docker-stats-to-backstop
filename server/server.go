// vim: ts=2:nowrap
package main

import (
	"flag"
	"github.com/deliverous/docker-stats-to-backstop/translate"
	"github.com/deliverous/docker-stats-to-backstop/translate/backstop"
	"github.com/deliverous/docker-stats-to-backstop/translate/docker"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	help        = flag.Bool("help", false, "Get help")
	backstopUrl = flag.String("backstop", env("SRV_BACKSTOP", ""), "URL for connecting backsop server")
	dockerUrl   = flag.String("docker", env("SRV_DOCKER", "unix:///var/run/docker.sock"), "URL for connecting docker server")
	prefix      = flag.String("prefix", "", "JSON containing 'regexp' and 'into' to rewrite the container name into graphite identifier")
	poll        = flag.String("poll", "5m", "Set the poll delay. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'")
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	prefixRule, err := loadPrefixRule(*prefix)
	if err != nil {
		log.Fatalf("ERROR: cannot load prefix rules: %s", err)
	}

	duration, err := time.ParseDuration(*poll)
	if err != nil {
		log.Fatalf("ERROR: cannot parse duration '%s': %s", *poll, err)
	}

	transport := &http.Transport{}
	transport.RegisterProtocol("unix", NewSocketTransport(LstatSocketPredicate, 2*time.Second))
	client := &http.Client{Transport: transport}

	dockerApi := docker.NewDockerApi(client, *dockerUrl)

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
				err = backstop.SendMetrics(client, *backstopUrl, translate.Translate(prefix, stats))
				if err != nil {
					log.Printf("ERROR: cannot send container stats: %s\n", err)
				}
			}
		}

		time.Sleep(duration)
	}
}

func env(key string, missing string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return missing
}
