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
	help        = flag.Bool("help", false, "Get help")
	backstopUrl = flag.String("backstop", env("SRV_BACKSTOP", ""), "URL for connecting backsop server")
	dockerUrl   = flag.String("docker", env("SRV_DOCKER", "unix:///var/run/docker.sock"), "URL for connecting docker server")
	prefix      = flag.String("prefix", env("SRV_PREFIX", ""), "JSON containing 'regexp' and 'into' to rewrite the container name into graphite identifier")
	poll        = flag.String("poll", env("SRV_PREFIX", "5m"), "Set the poll delay. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'")
	verbose     = flag.Bool("verbose", false, "Enable the verbose mode")
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	duration, err := time.ParseDuration(*poll)
	if err != nil {
		log.Fatalf("ERROR: cannot parse duration '%s': %s", *poll, err)
	}

	server.ServeForever(*dockerUrl, *backstopUrl, *prefix, duration, *verbose)
}

func env(key string, missing string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return missing
}
