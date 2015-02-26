// vim: ts=2 nowrap
package server

import (
	"errors"
	"net"
	"net/http"
	"net/url"
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
	socket, path, err := parseUnixUrl(*request.URL, transport.predicate)
	if err != nil {
		return nil, err
	}

	inner := &http.Transport{
		DisableCompression: true,
		Dial: func(proto, addr string) (conn net.Conn, err error) {
			return net.DialTimeout("unix", socket, transport.timeout)
		}}

	request.URL.Scheme = "http"
	request.URL.Host = socket
	request.URL.Path = path
	return inner.RoundTrip(request)
}

func parseUnixUrl(u url.URL, predicate SocketPredicate) (string, string, error) {
	urlPath := path.Join(u.Host, u.Path)
	if !path.IsAbs(urlPath) {
		urlPath = "/" + urlPath
	}
	socket, err := findSocket(urlPath, predicate)
	if err != nil {
		return "", "", err
	}
	return socket, strings.TrimPrefix(urlPath, socket), nil
}

func findSocket(urlPath string, predicate SocketPredicate) (string, error) {
	for urlPath != "/" {
		if predicate(urlPath) {
			return urlPath, nil
		}
		urlPath = path.Dir(urlPath)
	}
	return "", ErrNoSocket
}
