// vim: ts=2:nowrap
package server

import (
	"testing"
)

func Test_ReverseHostname(t *testing.T) {
	ensureReverseHostname(t, "hv-01", "hv-01")
	ensureReverseHostname(t, "hv-01.deliverous.net", "net.deliverous.hv-01")
}

func ensureReverseHostname(t *testing.T, hostname string, expected string) {
	reversed := reverseHostname(hostname)
	if reversed != expected {
		t.Errorf("reverse hostname failed:\nexpected %#v,\ngot      %#v", expected, reversed)
	}
}
