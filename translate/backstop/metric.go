// vim: ts=2:nowrap
package backstop

type Metric struct {
	Name      string `json:"metric"`
	Value     uint64 `json:"value"`
	Timestamp int64  `json:"measure_time"`
}
