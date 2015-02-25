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

	for _, container := range flag.Args() {
		stats, _ := docker.GetDockerStats(client, fmt.Sprintf("%s/v1.17/containers/%s/stats", *dockerUrl, container))
		metrics := translate.Translate("docker."+container, stats)
		backstop.SendMetrics(client, *backstopUrl, metrics)
	}
}
