// vim: ts=2 nowrap
package main

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"strings"
	"time"
)

var (
	ErrNoSocket = errors.New("No Unix domain socket found")
)

type SocketPredicate func(string) bool

func LstatSocketPredicate(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeType == os.ModeSocket
}

func ExtentionSocketPredicate(path string) bool {
	return strings.HasSuffix(path, ".sock")
}

func NewSocketTransport(predicate SocketPredicate, timeout time.Duration) http.RoundTripper {
	return &socketTransport{predicate: predicate, timeout: timeout}
}

type socketTransport struct {
	timeout   time.Duration
	predicate SocketPredicate
}

func (transport *socketTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	socket, err := findSocket(path.Join(request.URL.Host, request.URL.Path), transport.predicate)
	if err != nil {
		return nil, err
	}
	request.URL.Path = strings.TrimPrefix(request.URL.Path, socket)

	dial, err := net.DialTimeout("unix", socket, transport.timeout)
	if err != nil {
		return nil, err
	}
	socketClientConn := httputil.NewClientConn(dial, nil)
	defer socketClientConn.Close()
	return socketClientConn.Do(request)
}

func findSocket(urlPath string, predicate SocketPredicate) (string, error) {
	socket := path.Clean(urlPath)
	for socket != "" {
		if predicate(socket) {
			return socket, nil
		}
		socket = path.Dir(socket)
	}
	return "", ErrNoSocket
}
