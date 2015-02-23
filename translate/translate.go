// vim: ts=2 nowrap
package translate

import (
	"encoding/json"
)

type ContainerStats struct {
	Network *NetworkStats
}

type NetworkStats struct {
	RxBytes   int64 `json:"rx_bytes"`
	RxPackets int64 `json:"rx_packets"`
	RxErrors  int64 `json:"rx_errors"`
	RxDropped int64 `json:"rx_dropped"`
	TxBytes   int64 `json:"tx_bytes"`
	TxPackets int64 `json:"tx_packets"`
	TxErrors  int64 `json:"tx_errors"`
	TxDropped int64 `json:"tx_dropped"`
}

func ParseJson(data string) (*ContainerStats, error) {
	var content *ContainerStats
	err := json.Unmarshal([]byte(data), &content)
	return content, err
}

// Given : {"read":"2015-02-23T09:39:06.071088072+01:00","network":{"rx_bytes":120062,"rx_packets":1425,"rx_errors":0,"rx_dropped":0,"tx_bytes":160652,"tx_packets":1257,"tx_errors":0,"tx_dropped":0},"cpu_stats":{"cpu_usage":{"total_usage":817842610809,"percpu_usage":[171131425370,170731715729,171006335590,171746607003,31328388915,34118247352,36020588573,31759302277],"usage_in_kernelmode":44460000000,"usage_in_usermode":100440000000},"system_cpu_usage":6616189560000000,"throttling_data":{"periods":0,"throttled_periods":0,"throttled_time":0}},"memory_stats":{"usage":361316352,"max_usage":600657920,"stats":{"active_anon":354676736,"active_file":6475776,"cache":6717440,"hierarchical_memory_limit":2147483648,"hierarchical_memsw_limit":4294967296,"inactive_anon":0,"inactive_file":163840,"mapped_file":32768,"pgfault":710237,"pgmajfault":0,"pgpgin":669598,"pgpgout":581386,"rss":354598912,"rss_huge":0,"swap":0,"total_active_anon":354676736,"total_active_file":6475776,"total_cache":6717440,"total_inactive_anon":0,"total_inactive_file":163840,"total_mapped_file":32768,"total_pgfault":710237,"total_pgmajfault":0,"total_pgpgin":669598,"total_pgpgout":581386,"total_rss":354598912,"total_rss_huge":0,"total_swap":0,"total_unevictable":0,"total_writeback":0,"unevictable":0,"writeback":0},"failcnt":0,"limit":2147483648},"blkio_stats":{"io_service_bytes_recursive":[{"major":8,"minor":16,"op":"Read","value":0},{"major":8,"minor":16,"op":"Write","value":0},{"major":8,"minor":16,"op":"Sync","value":0},{"major":8,"minor":16,"op":"Async","value":0},{"major":8,"minor":16,"op":"Total","value":0},{"major":8,"minor":0,"op":"Read","value":0},{"major":8,"minor":0,"op":"Write","value":0},{"major":8,"minor":0,"op":"Sync","value":0},{"major":8,"minor":0,"op":"Async","value":0},{"major":8,"minor":0,"op":"Total","value":0}],"io_serviced_recursive":[{"major":8,"minor":16,"op":"Read","value":0},{"major":8,"minor":16,"op":"Write","value":0},{"major":8,"minor":16,"op":"Sync","value":0},{"major":8,"minor":16,"op":"Async","value":0},{"major":8,"minor":16,"op":"Total","value":0},{"major":8,"minor":0,"op":"Read","value":0},{"major":8,"minor":0,"op":"Write","value":0},{"major":8,"minor":0,"op":"Sync","value":0},{"major":8,"minor":0,"op":"Async","value":0},{"major":8,"minor":0,"op":"Total","value":0}],"io_queue_recursive":[{"major":8,"minor":16,"op":"Read","value":0},{"major":8,"minor":16,"op":"Write","value":0},{"major":8,"minor":16,"op":"Sync","value":0},{"major":8,"minor":16,"op":"Async","value":0},{"major":8,"minor":16,"op":"Total","value":0},{"major":8,"minor":0,"op":"Read","value":0},{"major":8,"minor":0,"op":"Write","value":0},{"major":8,"minor":0,"op":"Sync","value":0},{"major":8,"minor":0,"op":"Async","value":0},{"major":8,"minor":0,"op":"Total","value":0}],"io_service_time_recursive":[{"major":8,"minor":16,"op":"Read","value":0},{"major":8,"minor":16,"op":"Write","value":0},{"major":8,"minor":16,"op":"Sync","value":0},{"major":8,"minor":16,"op":"Async","value":0},{"major":8,"minor":16,"op":"Total","value":0},{"major":8,"minor":0,"op":"Read","value":0},{"major":8,"minor":0,"op":"Write","value":0},{"major":8,"minor":0,"op":"Sync","value":0},{"major":8,"minor":0,"op":"Async","value":0},{"major":8,"minor":0,"op":"Total","value":0}],"io_wait_time_recursive":[{"major":8,"minor":16,"op":"Read","value":0},{"major":8,"minor":16,"op":"Write","value":0},{"major":8,"minor":16,"op":"Sync","value":0},{"major":8,"minor":16,"op":"Async","value":0},{"major":8,"minor":16,"op":"Total","value":0},{"major":8,"minor":0,"op":"Read","value":0},{"major":8,"minor":0,"op":"Write","value":0},{"major":8,"minor":0,"op":"Sync","value":0},{"major":8,"minor":0,"op":"Async","value":0},{"major":8,"minor":0,"op":"Total","value":0}],"io_merged_recursive":[{"major":8,"minor":16,"op":"Read","value":0},{"major":8,"minor":16,"op":"Write","value":0},{"major":8,"minor":16,"op":"Sync","value":0},{"major":8,"minor":16,"op":"Async","value":0},{"major":8,"minor":16,"op":"Total","value":0},{"major":8,"minor":0,"op":"Read","value":0},{"major":8,"minor":0,"op":"Write","value":0},{"major":8,"minor":0,"op":"Sync","value":0},{"major":8,"minor":0,"op":"Async","value":0},{"major":8,"minor":0,"op":"Total","value":0}],"io_time_recursive":[{"major":8,"minor":16,"op":"","value":0},{"major":8,"minor":0,"op":"","value":0}],"sectors_recursive":[{"major":8,"minor":16,"op":"","value":0},{"major":8,"minor":0,"op":"","value":0}]}}
