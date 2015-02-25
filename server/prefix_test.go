// vim: ts=2:nowrap
package server

import (
	"github.com/deliverous/docker-stats-to-backstop/translate/docker"
	"testing"
)

func Test_ComputePrefix(t *testing.T) {
	ensurePrefix(t, "", container("id"), "id")
	ensurePrefix(t, "", container("id", "the name"), "the name")
	ensurePrefix(t, `{"regexp": "(.*)\\..*", "into": "prefix.$1"}`, container("id", "name.bla"), "prefix.name")
	ensurePrefix(t, `{"regexp": "(.*)\\..*\\.(.*)\\..*", "into": "$1.$2"}`, container("id", "pid.bla.name.service"), "pid.name")
}

func ensurePrefix(t *testing.T, definition string, container docker.Container, expected string) {
	rule, err := loadPrefixRule(definition)
	if err != nil {
		t.Errorf("unexpected error on '%s': %#v", definition, err)
	}
	prefix := computePrefix(&container, rule)
	if prefix != expected {
		t.Errorf("rewrite failed:\nexpected %#v,\ngot      %#v", expected, prefix)
	}
}

func container(id string, names ...string) docker.Container {
	return docker.Container{Id: "id", Names: names}
}
