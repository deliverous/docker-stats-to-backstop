// vim: ts=2:nowrap
package main

import (
	"flag"
	"fmt"
	"github.com/deliverous/docker-stats-to-backstop/translate"
	"github.com/deliverous/docker-stats-to-backstop/translate/backstop"
	"github.com/deliverous/docker-stats-to-backstop/translate/docker"
	"gopkg.in/jmcvetta/napping.v1"
	"net/http"
	"net/url"
	"os"
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

func getDockerStats(urlString string) (*docker.ContainerStats, error) {
	transport := &http.Transport{}
	transport.RegisterProtocol("unix", NewSocketTransport())
	client := &http.Client{Transport: transport}

	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	fmt.Printf("url: %#v\n", u)
	session := &napping.Session{Userinfo: u.User, Log: false, Client: client}
	stats := &docker.ContainerStats{}
	u.User = nil
	_, err = docker.GetDockerStats(session, u.String(), stats)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func main() {
	flag.Parse()
	fmt.Printf("== config =======================\n")
	fmt.Printf("docker URL  : %s\n", *dockerUrl)
	fmt.Printf("backstop URL: %s\n", *backstopUrl)
	fmt.Printf("== starting =====================\n")

	for _, container := range flag.Args() {
		stats, _ := getDockerStats(fmt.Sprintf("%s/v1.17/containers/%s/stats", *dockerUrl, container))
		metrics := translate.Translate("docker."+container, stats)
		data, _ := backstop.RenderMetrics(metrics)
		fmt.Printf("Metrics: %s\n", data)
	}
}
