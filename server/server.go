// vim: ts=2:nowrap
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/deliverous/docker-stats-to-backstop/translate"
	"github.com/deliverous/docker-stats-to-backstop/translate/backstop"
	"github.com/deliverous/docker-stats-to-backstop/translate/docker"
	"log"
	"net/http"
	"os"
	"regexp"
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
	prefix      = flag.String("prefix", "", "JSON containing 'regexp' and 'into' to rewrite the container name into graphite identifier")
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	prefixRule := loadPrefixRule(*prefix)

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
		prefix := computePrefix(&container, prefixRule)

		fmt.Printf("Processing container %s (%s)\n", container.Id, prefix)
		stats, err := dockerApi.GetContainerStats(container.Id)
		if err != nil {
			fmt.Printf("ERROR: cannot get container stats: %s\n", err)
			continue
		}
		err = backstop.SendMetrics(client, *backstopUrl, translate.Translate(prefix, stats))
		if err != nil {
			fmt.Printf("ERROR: cannot send container stats: %s\n", err)
		}
	}
}

type prefixRule struct {
	Regexp string
	Into   string
	parsed *regexp.Regexp
}

func loadPrefixRule(definition string) *prefixRule {
	if definition == "" {
		return nil
	}

	s := prefixRule{}
	if err := json.Unmarshal([]byte(definition), &s); err != nil {
		log.Fatalf("ERROR: invalid rewrite definition: %s", err)
	}
	if r, err := regexp.Compile(s.Regexp); err != nil {
		log.Fatalf("ERROR: invalid regexp: %s", err)
	} else {
		s.parsed = r
	}
	return &s
}

func computePrefix(container *docker.Container, rule *prefixRule) string {
	prefix := container.Id
	if len(container.Names) > 0 {
		prefix = container.Names[0]
		if rule != nil {
			return rule.parsed.ReplaceAllString(prefix, rule.Into)
		}
	}
	return prefix
}
