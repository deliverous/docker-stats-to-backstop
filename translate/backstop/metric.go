// vim: ts=2:nowrap
package backstop

import (
	"time"
)

type Metric struct {
	Name      string    `json:"metric"`
	Value     int64     `json:"value"`
	Timestamp time.Time `json:"measure_time"`
}
