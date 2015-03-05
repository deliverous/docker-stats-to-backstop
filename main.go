// vim: ts=2:nowrap
package main

import (
	"flag"
	"github.com/deliverous/docker-stats-to-backstop/server"
	"log"
	"os"
	"time"
)

var (
	help            = flag.Bool("help", false, "Get help")
	backstopURL     = flag.String("backstop", env("SRV_BACKSTOP", ""), "URL for connecting backsop server")
	dockerURL       = flag.String("docker", env("SRV_DOCKER", "unix:///var/run/docker.sock"), "URL for connecting docker server")
	rulesDefinition = flag.String("rules", env("SRV_RULES", ""), "JSON containing 'regexp','into' and 'category' to rewrite the container name into graphite identifier")
	pollDefinition  = flag.String("poll", env("SRV_POLL", "5m"), "Set the poll delay. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'")
	verbose         = flag.Bool("verbose", false, "Enable the verbose mode")
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	poll, err := time.ParseDuration(*pollDefinition)
	if err != nil {
		log.Fatalf("ERROR: cannot parse duration '%s': %s", *pollDefinition, err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("ERROR: cannot get hostname : %s", err)
	}

	rules, err := server.LoadRules(*rulesDefinition)
	if err != nil {
		log.Fatalf("ERROR: cannot load rules '%s' : %s", *rulesDefinition, err)
	}

	srv := &server.Server{
		Hostname:    hostname,
		DockerURL:   *dockerURL,
		BackstopURL: *backstopURL,
		Poll:        poll,
		Rules:       rules,
		Verbose:     *verbose,
	}

	srv.Serve()
}

func env(key string, missing string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return missing
}
