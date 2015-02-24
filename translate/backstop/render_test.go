// vim: ts=2:nowrap
package backstop

import (
	"testing"
	"time"
)

func Test_RenderMetrics(t *testing.T) {
	now := time.Date(2015, time.January, 23, 15, 39, 06, 00, time.UTC)
	cpu := Metric{Name: "cpu", Value: 42, Timestamp: now}
	memory := Metric{Name: "memory.enough", Value: 512, Timestamp: now}
	data, _ := RenderMetrics([]Metric{cpu, memory})

	if data != `[{"metric":"cpu","value":42,"measure_time":"2015-01-23T15:39:06Z"},{"metric":"memory.enough","value":512,"measure_time":"2015-01-23T15:39:06Z"}]` {
		t.Errorf("render metrics failed: got '%s'", data)
	}
}
