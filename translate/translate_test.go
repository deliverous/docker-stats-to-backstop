// vim: ts=2:nowrap
package translate

import (
	"testing"
)

func Test_ParseJson_NetworkStats_Rx(t *testing.T) {
	stats, _ := ParseJson(`{"network":{"rx_bytes":120062,"rx_packets":1425,"rx_errors":1,"rx_dropped":2}}`)
	checkIntValue(t, "RxBytes", 120062, stats.Network.RxBytes)
	checkIntValue(t, "RxPackets", 1425, stats.Network.RxPackets)
	checkIntValue(t, "RxErrors", 1, stats.Network.RxErrors)
	checkIntValue(t, "RxDropped", 2, stats.Network.RxDropped)
}

func Test_ParseJson_NetworkStats_Tx(t *testing.T) {
	stats, _ := ParseJson(`{"network":{"tx_bytes":160652,"tx_packets":1257,"tx_errors":3,"tx_dropped":5}}`)
	checkIntValue(t, "TxBytes", 160652, stats.Network.TxBytes)
	checkIntValue(t, "TxPackets", 1257, stats.Network.TxPackets)
	checkIntValue(t, "TxErrors", 3, stats.Network.TxErrors)
	checkIntValue(t, "TxDropped", 5, stats.Network.TxDropped)
}

func Test_ParseJson_CpuStats(t *testing.T) {
	stats, _ := ParseJson(`{"cpu_stats":{"cpu_usage":{"total_usage":817842610809,"usage_in_kernelmode":44460000000,"usage_in_usermode":100440000000},"system_cpu_usage":6616189560000000}}`)

	checkIntValue(t, "CupUsage.TotalUsage", 817842610809, stats.Cpu.CpuUsage.TotalUsage)
	checkIntValue(t, "CupUsage.UsageInKernelmode", 44460000000, stats.Cpu.CpuUsage.UsageInKernelmode)
	checkIntValue(t, "CupUsage.UsageInUsermode", 100440000000, stats.Cpu.CpuUsage.UsageInUsermode)
	checkIntValue(t, "SystemCpuUsage", 6616189560000000, stats.Cpu.SystemCpuUsage)
}

func Test_ParseJson_MemoryStats(t *testing.T) {
	stats, _ := ParseJson(`{"memory_stats":{ "usage":361316352, "max_usage":600657920, "limit":2147483648}}`)

	checkIntValue(t, "Usage", 361316352, stats.Memory.Usage)
	checkIntValue(t, "MaxUsage", 600657920, stats.Memory.MaxUsage)
	checkIntValue(t, "Limit", 2147483648, stats.Memory.Limit)
}
func checkIntValue(t *testing.T, fieldName string, expected int64, got int64) {
	if got != expected {
		t.Errorf("'%s' verification failed: got %d, expected %d", fieldName, got, expected)
	}
}

const jsonContainerStats = `
  {
    "read":"2015-02-23T09:39:06.071088072+01:00",
    "network":{"rx_bytes":120062,"rx_packets":1425,"rx_errors":1,"rx_dropped":2,"tx_bytes":160652,"tx_packets":1257,"tx_errors":3,"tx_dropped":5},
    "cpu_stats":{
      "cpu_usage":{
        "total_usage":817842610809,
        "percpu_usage":[171131425370,170731715729,171006335590,171746607003,31328388915,34118247352,36020588573,31759302277],
        "usage_in_kernelmode":44460000000,
        "usage_in_usermode":100440000000},
      "system_cpu_usage":6616189560000000,
      "throttling_data":{"periods":0,"throttled_periods":0,"throttled_time":0}},
    "memory_stats":{
      "usage":361316352,
      "max_usage":600657920,
      "stats":{"active_anon":354676736,"active_file":6475776,"cache":6717440,"hierarchical_memory_limit":2147483648,"hierarchical_memsw_limit":4294967296,"inactive_anon":0,"inactive_file":163840,"mapped_file":32768,"pgfault":710237,"pgmajfault":0,"pgpgin":669598,"pgpgout":581386,"rss":354598912,"rss_huge":0,"swap":0,"total_active_anon":354676736,"total_active_file":6475776,"total_cache":6717440,"total_inactive_anon":0,"total_inactive_file":163840,"total_mapped_file":32768,"total_pgfault":710237,"total_pgmajfault":0,"total_pgpgin":669598,"total_pgpgout":581386,"total_rss":354598912,"total_rss_huge":0,"total_swap":0,"total_unevictable":0,"total_writeback":0,"unevictable":0,"writeback":0},
      "failcnt":0,
      "limit":2147483648},
    "blkio_stats":{"io_service_bytes_recursive":[{"major":8,"minor":16,"op":"Read","value":0},{"major":8,"minor":16,"op":"Write","value":0},{"major":8,"minor":16,"op":"Sync","value":0},{"major":8,"minor":16,"op":"Async","value":0},{"major":8,"minor":16,"op":"Total","value":0},{"major":8,"minor":0,"op":"Read","value":0},{"major":8,"minor":0,"op":"Write","value":0},{"major":8,"minor":0,"op":"Sync","value":0},{"major":8,"minor":0,"op":"Async","value":0},{"major":8,"minor":0,"op":"Total","value":0}],"io_serviced_recursive":[{"major":8,"minor":16,"op":"Read","value":0},{"major":8,"minor":16,"op":"Write","value":0},{"major":8,"minor":16,"op":"Sync","value":0},{"major":8,"minor":16,"op":"Async","value":0},{"major":8,"minor":16,"op":"Total","value":0},{"major":8,"minor":0,"op":"Read","value":0},{"major":8,"minor":0,"op":"Write","value":0},{"major":8,"minor":0,"op":"Sync","value":0},{"major":8,"minor":0,"op":"Async","value":0},{"major":8,"minor":0,"op":"Total","value":0}],"io_queue_recursive":[{"major":8,"minor":16,"op":"Read","value":0},{"major":8,"minor":16,"op":"Write","value":0},{"major":8,"minor":16,"op":"Sync","value":0},{"major":8,"minor":16,"op":"Async","value":0},{"major":8,"minor":16,"op":"Total","value":0},{"major":8,"minor":0,"op":"Read","value":0},{"major":8,"minor":0,"op":"Write","value":0},{"major":8,"minor":0,"op":"Sync","value":0},{"major":8,"minor":0,"op":"Async","value":0},{"major":8,"minor":0,"op":"Total","value":0}],"io_service_time_recursive":[{"major":8,"minor":16,"op":"Read","value":0},{"major":8,"minor":16,"op":"Write","value":0},{"major":8,"minor":16,"op":"Sync","value":0},{"major":8,"minor":16,"op":"Async","value":0},{"major":8,"minor":16,"op":"Total","value":0},{"major":8,"minor":0,"op":"Read","value":0},{"major":8,"minor":0,"op":"Write","value":0},{"major":8,"minor":0,"op":"Sync","value":0},{"major":8,"minor":0,"op":"Async","value":0},{"major":8,"minor":0,"op":"Total","value":0}],"io_wait_time_recursive":[{"major":8,"minor":16,"op":"Read","value":0},{"major":8,"minor":16,"op":"Write","value":0},{"major":8,"minor":16,"op":"Sync","value":0},{"major":8,"minor":16,"op":"Async","value":0},{"major":8,"minor":16,"op":"Total","value":0},{"major":8,"minor":0,"op":"Read","value":0},{"major":8,"minor":0,"op":"Write","value":0},{"major":8,"minor":0,"op":"Sync","value":0},{"major":8,"minor":0,"op":"Async","value":0},{"major":8,"minor":0,"op":"Total","value":0}],"io_merged_recursive":[{"major":8,"minor":16,"op":"Read","value":0},{"major":8,"minor":16,"op":"Write","value":0},{"major":8,"minor":16,"op":"Sync","value":0},{"major":8,"minor":16,"op":"Async","value":0},{"major":8,"minor":16,"op":"Total","value":0},{"major":8,"minor":0,"op":"Read","value":0},{"major":8,"minor":0,"op":"Write","value":0},{"major":8,"minor":0,"op":"Sync","value":0},{"major":8,"minor":0,"op":"Async","value":0},{"major":8,"minor":0,"op":"Total","value":0}],"io_time_recursive":[{"major":8,"minor":16,"op":"","value":0},{"major":8,"minor":0,"op":"","value":0}],"sectors_recursive":[{"major":8,"minor":16,"op":"","value":0},{"major":8,"minor":0,"op":"","value":0}]}
  }
`
