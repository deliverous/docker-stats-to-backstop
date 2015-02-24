// vim: ts=2:nowrap
package main

import (
	"flag"
	"fmt"
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
	dockerUrl   = flag.String("docker", env("SRV_DOCKER", "unix:///var/run/docker.sock/v1.17/info"), "URL for connecting docker server")
)

func getDockerStats(urlString string) error {
	transport := &http.Transport{}
	transport.RegisterProtocol("unix", NewSocketTransport())
	client := &http.Client{Transport: transport}

	u, err := url.Parse(urlString)
	if err != nil {
		return err
	}
	fmt.Printf("url: %#v\n", u)
	session := &napping.Session{Userinfo: u.User, Log: true, Client: client}
	stats := &docker.ContainerStats{}
	u.User = nil
	docker.GetDockerStats(session, u.String(), stats)
	return nil
}

func main() {
	flag.Parse()
	fmt.Printf("== config =======================\n")
	fmt.Printf("docker URL  : %s\n", *dockerUrl)
	fmt.Printf("backstop URL: %s\n", *backstopUrl)
	fmt.Printf("== starting =====================\n")

	getDockerStats(*dockerUrl)
	//dockerSession, err := buildSession(*dockerUrl)
	//panicOnError(err)
	//stats := &docker.ContainerStats{}
	//docker.GetDockerStats(dockerSession, "", stats)
	//backstopSession, err := buildSession(backstop)
	//panicOnError(err)

}
