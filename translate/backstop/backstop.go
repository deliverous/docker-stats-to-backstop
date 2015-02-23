// vim: ts=2:nowrap
package backstop

import (
	"encoding/json"
	"time"
)

type Metric struct {
	Name  string    `json:"metric"`
	Value int64     `json:"value"`
	Time  time.Time `json:"measure_time"`
}

type Metrics []Metric

func RenderMetrics(metrics []Metric) (string, error) {
	data, err := json.Marshal(metrics)
	return string(data), err
}
