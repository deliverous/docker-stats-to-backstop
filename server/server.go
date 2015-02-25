// vim: ts=2:nowrap
package main

import (
	"flag"
	"fmt"
	"github.com/deliverous/docker-stats-to-backstop/translate"
	"github.com/deliverous/docker-stats-to-backstop/translate/backstop"
	"github.com/deliverous/docker-stats-to-backstop/translate/docker"
	"net/http"
	"os"
	"time"
)

func env(key string, missing string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return missing
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	help        = flag.Bool("help", false, "Get help")
	backstopUrl = flag.String("backstop", env("SRV_BACKSTOP", ""), "URL for connecting backsop server")
	dockerUrl   = flag.String("docker", env("SRV_DOCKER", "unix:///var/run/docker.sock"), "URL for connecting docker server")
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	transport := &http.Transport{}
	transport.RegisterProtocol("unix", NewSocketTransport(LstatSocketPredicate, 2*time.Second))
	client := &http.Client{Transport: transport}

	dockerApi := docker.NewDockerApi(client, *dockerUrl)

	version, err := dockerApi.GetApiVersion()
	panicOnError(err)
	if version != docker.ApiVersion {
		fmt.Println("WARNING: not using the latest api version")
		fmt.Printf("Using  : '%s'\n", docker.ApiVersion)
		fmt.Printf("Current: '%s'\n", version)
	}

	containers, err := dockerApi.GetContainers()
	panicOnError(err)

	for _, container := range containers {
		fmt.Printf("Processing container %s\n", container.Id)
		stats, _ := dockerApi.GetContainerStats(container.Id)
		metrics := translate.Translate("docker."+container.Id, stats)
		fmt.Printf("metrics: %#v\n", metrics)
		backstop.SendMetrics(client, *backstopUrl, metrics)
	}
}
