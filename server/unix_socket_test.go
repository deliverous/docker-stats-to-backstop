// vim: ts=2:nowrap
package server

import (
	"net/url"
	"testing"
)

func Test_ParseUnixUrl(t *testing.T) {
	checkSuccess(t, url.URL{Path: "/var/run/docker.sock"}, "/var/run/docker.sock", "")
	checkSuccess(t, url.URL{Path: "/var/run/docker.sock/"}, "/var/run/docker.sock", "")
	checkSuccess(t, url.URL{Path: "/var/run/docker.sock/info"}, "/var/run/docker.sock", "/info")
	checkSuccess(t, url.URL{Host: "var", Path: "/run/docker.sock/info"}, "/var/run/docker.sock", "/info")
	checkError(t, url.URL{Path: "/no/socket"}, ErrNoSocket)
	checkError(t, url.URL{Host: "var", Path: "/no/socket"}, ErrNoSocket)
}

func checkSuccess(t *testing.T, input url.URL, expectedSocket, expectedPath string) {
	socket, path, err := parseUnixUrl(input, ExtentionSocketPredicate)
	if err != nil {
		t.Errorf("parse unix url failed:\ncase %#v,\nunexpected error %#v", input, err)
	} else {
		if socket != expectedSocket {
			t.Errorf("parse unix url failed:\ncase     %#v,\nexpected %#v,\ngot      %#v", input, expectedSocket, socket)
		}
		if path != expectedPath {
			t.Errorf("parse unix url failed:\ncase     %#v,\nexpected %#v,\ngot      %#v", input, expectedPath, path)
		}
	}
}

func checkError(t *testing.T, input url.URL, expected error) {
	_, _, err := parseUnixUrl(input, ExtentionSocketPredicate)
	if err != expected {
		t.Errorf("update url failed:\ncase           %#v,\nexpected error %#v,\ngot            %#v", input, expected, err)
	}
}
