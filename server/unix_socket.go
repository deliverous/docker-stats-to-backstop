package main

import (
	"bufio"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"strings"
	"time"
)

func NewSocketTransport() http.RoundTripper {
	return &socketTransport{}
}

type socketTransport struct {
}

func (transport *socketTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	socket, err := split(request.URL.Path)
	if err != nil {
		return nil, err
	}
	request.URL.Path = strings.TrimPrefix(request.URL.Path, socket)

	dial, err := net.DialTimeout("unix", socket, 32*time.Second)
	if err != nil {
		return nil, err
	}
	socketClientConn := httputil.NewClientConn(dial, nil)
	defer socketClientConn.Close()

	response, err := socketClientConn.Do(request)

	// Hack {
	line, err := bufio.NewReader(response.Body).ReadString('\n')
	if err != nil {
		return nil, err
	}
	response.Body = ioutil.NopCloser(strings.NewReader(line))
	// Hack }

	return response, err
}

func split(urlPath string) (string, error) {
	socket := path.Clean(urlPath)
	for socket != "" {
		if isUnixSocket(socket) {
			return socket, nil
		}
		socket = path.Dir(socket)
	}
	return "", errors.New("No Unix domain socket found in '" + urlPath + "'")
}

func isUnixSocket(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeType == os.ModeSocket
}
