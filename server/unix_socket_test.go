// vim: ts=2:nowrap
package main

import (
	"net/url"
	"testing"
)

func Test_UpdateUrl(t *testing.T) {
	checkSuccess(t, url.URL{Path: "/var/run/docker.sock"}, url.URL{Host: "/var/run/docker.sock"})
	checkSuccess(t, url.URL{Path: "/var/run/docker.sock/"}, url.URL{Host: "/var/run/docker.sock", Path: "/"})
	checkSuccess(t, url.URL{Path: "/var/run/docker.sock/info"}, url.URL{Host: "/var/run/docker.sock", Path: "/info"})
	checkSuccess(t, url.URL{Host: "var", Path: "/run/docker.sock/info"}, url.URL{Host: "/var/run/docker.sock", Path: "/info"})
	checkError(t, url.URL{Path: "/no/socket"}, ErrNoSocket)
	checkError(t, url.URL{Host: "var", Path: "/no/socket"}, ErrNoSocket)
}

func checkSuccess(t *testing.T, input, expected url.URL) {
	u, err := updateUrl(input, ExtentionSocketPredicate)
	if err != nil {
		t.Errorf("update url failed:\ncase %#v,\nunexpected error %#v", input, err)
	} else if u != expected {
		t.Errorf("update url failed:\ncase     %#v,\nexpected %#v,\ngot      %#v", input, expected, u)
	}
}

func checkError(t *testing.T, input url.URL, expected error) {
	_, err := updateUrl(input, ExtentionSocketPredicate)
	if err != expected {
		t.Errorf("update url failed:\ncase           %#v,\nexpected error %#v,\ngot            %#v", input, expected, err)
	}
}
