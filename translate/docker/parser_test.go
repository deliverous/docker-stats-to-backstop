// vim: ts=2:nowrap
package docker

import (
	"testing"
	"time"
)

func Test_ParseDockerStats_Timestamp(t *testing.T) {
	stats, _ := ParseDockerStats(`{"read":"2015-02-23T09:39:06.071088072+01:00"}`)
	expected := time.Date(2015, time.February, 23, 8, 39, 6, 0, time.UTC)
	if stats.Timestamp.Unix() != expected.Unix() {
		t.Errorf("timestamp: got: %s, expected: %s", stats.Timestamp.UTC().String(), expected.String())
	}
}

func Test_ParseDockerStats_NetworkStats_Rx(t *testing.T) {
	stats, _ := ParseDockerStats(`{"network":{"rx_bytes":120062,"rx_packets":1425,"rx_errors":1,"rx_dropped":2}}`)
	checkIntValue(t, "RxBytes", 120062, stats.Network.RxBytes)
	checkIntValue(t, "RxPackets", 1425, stats.Network.RxPackets)
	checkIntValue(t, "RxErrors", 1, stats.Network.RxErrors)
	checkIntValue(t, "RxDropped", 2, stats.Network.RxDropped)
}

func Test_ParseDockerStats_NetworkStats_Tx(t *testing.T) {
	stats, _ := ParseDockerStats(`{"network":{"tx_bytes":160652,"tx_packets":1257,"tx_errors":3,"tx_dropped":5}}`)
	checkIntValue(t, "TxBytes", 160652, stats.Network.TxBytes)
	checkIntValue(t, "TxPackets", 1257, stats.Network.TxPackets)
	checkIntValue(t, "TxErrors", 3, stats.Network.TxErrors)
	checkIntValue(t, "TxDropped", 5, stats.Network.TxDropped)
}

func Test_ParseDockerStats_CpuStats(t *testing.T) {
	stats, _ := ParseDockerStats(`{"cpu_stats":{"cpu_usage":{"total_usage":817842610809,"usage_in_kernelmode":44460000000,"usage_in_usermode":100440000000},"system_cpu_usage":6616189560000000}}`)

	checkIntValue(t, "CupUsage.TotalUsage", 817842610809, stats.Cpu.CpuUsage.TotalUsage)
	checkIntValue(t, "CupUsage.UsageInKernelmode", 44460000000, stats.Cpu.CpuUsage.UsageInKernelmode)
	checkIntValue(t, "CupUsage.UsageInUsermode", 100440000000, stats.Cpu.CpuUsage.UsageInUsermode)
	checkIntValue(t, "SystemCpuUsage", 6616189560000000, stats.Cpu.SystemCpuUsage)
}

func Test_ParseDockerStats_MemoryStats(t *testing.T) {
	stats, _ := ParseDockerStats(`{"memory_stats":{ "usage":361316352, "max_usage":600657920, "limit":2147483648}}`)

	checkIntValue(t, "Usage", 361316352, stats.Memory.Usage)
	checkIntValue(t, "MaxUsage", 600657920, stats.Memory.MaxUsage)
	checkIntValue(t, "Limit", 2147483648, stats.Memory.Limit)
}

func checkIntValue(t *testing.T, fieldName string, expected uint64, got *uint64) {
	if *got != expected {
		t.Errorf("'%s' verification failed: got %d, expected %d", fieldName, *got, expected)
	}
}
