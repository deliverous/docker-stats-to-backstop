// vim: ts=2:nowrap
package docker

import (
	"encoding/json"
	"testing"
	"time"
)

func Test_ParseDockerJson_Name(t *testing.T) {
	json, _ := parseDockerJson(`{"Name": "/1e70b2e7-55e5-403e-95ca-1e525c2bcee8.skillreso.rails.service"}`)
	expected := "/1e70b2e7-55e5-403e-95ca-1e525c2bcee8.skillreso.rails.service"
	if json.Name != expected {
		t.Errorf("name: got: %s, expected: %s", json.Name, expected)
	}
}

func Test_ParseDockerJson_StartedAt(t *testing.T) {
	json, _ := parseDockerJson(`{"State": { "StartedAt": "2015-05-12T10:27:19.213199747Z" }}`)
	expected := time.Date(2015, time.May, 12, 10, 27, 19, 0, time.UTC)
	if json.State.StartedAt.Unix() != expected.Unix() {
		t.Errorf("started at: got: %s, expected: %s", json.State.StartedAt.UTC().String(), expected.String())
	}
}

func parseDockerJson(data string) (*ContainerJson, error) {
	var content *ContainerJson
	err := json.Unmarshal([]byte(data), &content)
	return content, err
}
