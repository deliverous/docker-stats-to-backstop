// vim: ts=2:nowrap
package backstop

import (
	"encoding/json"
)

func RenderMetrics(metrics []Metric) (string, error) {
	data, err := json.Marshal(metrics)
	return string(data), err
}
