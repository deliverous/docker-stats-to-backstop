// vim: ts=2 nowrap
package server

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
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
	u, err := updateUrl(*request.URL, transport.predicate)
	if err != nil {
		return nil, err
	}
	request.URL = &u

	dial, err := net.DialTimeout("unix", u.Host, transport.timeout)
	if err != nil {
		return nil, err
	}
	socketClientConn := httputil.NewClientConn(dial, nil)
	defer socketClientConn.Close()
	return socketClientConn.Do(request)
}

func updateUrl(u url.URL, predicate SocketPredicate) (url.URL, error) {
	urlPath := path.Join(u.Host, u.Path)
	if !path.IsAbs(urlPath) {
		urlPath = "/" + urlPath
	}
	socket, err := findSocket(urlPath, predicate)
	if err != nil {
		return u, err
	}
	u.Host = socket
	u.Path = strings.TrimPrefix(urlPath, socket)
	return u, nil
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
