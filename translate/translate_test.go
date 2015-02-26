// vim: ts=2:nowrap
package translate

import (
	"github.com/deliverous/docker-stats-to-backstop/translate/docker"
	"testing"
	"time"
)

func Test_CpuStats(t *testing.T) {
	checkCpuStats(t, docker.CpuStats{SystemCpuUsage: v(100)}, "prefix.cpu.system", 100)
	checkCpuStats(t, docker.CpuStats{CpuUsage: docker.CpuUsageStats{TotalUsage: v(3000)}}, "prefix.cpu.total", 3000)
	checkCpuStats(t, docker.CpuStats{CpuUsage: docker.CpuUsageStats{UsageInKernelmode: v(2000)}}, "prefix.cpu.kernel", 2000)
	checkCpuStats(t, docker.CpuStats{CpuUsage: docker.CpuUsageStats{UsageInUsermode: v(1000)}}, "prefix.cpu.user", 1000)
}

func checkCpuStats(t *testing.T, stats docker.CpuStats, name string, value int64) {
	checkTranslation(t, &docker.ContainerStats{Cpu: stats}, name, value)
}

func Test_MemoryStats(t *testing.T) {
	checkMemoryStats(t, docker.MemoryStats{Usage: v(100)}, "prefix.memory.usage", 100)
	checkMemoryStats(t, docker.MemoryStats{MaxUsage: v(200)}, "prefix.memory.max_usage", 200)
	checkMemoryStats(t, docker.MemoryStats{Limit: v(512)}, "prefix.memory.limit", 512)
}

func checkMemoryStats(t *testing.T, stats docker.MemoryStats, name string, value int64) {
	checkTranslation(t, &docker.ContainerStats{Memory: stats}, name, value)
}

func Test_NetworkStats(t *testing.T) {
	checkNetworkStats(t, docker.NetworkStats{RxBytes: v(100)}, "prefix.network.rx_bytes", 100)
	checkNetworkStats(t, docker.NetworkStats{RxPackets: v(100)}, "prefix.network.rx_packets", 100)
	checkNetworkStats(t, docker.NetworkStats{RxErrors: v(100)}, "prefix.network.rx_errors", 100)
	checkNetworkStats(t, docker.NetworkStats{RxDropped: v(100)}, "prefix.network.rx_dropped", 100)
	checkNetworkStats(t, docker.NetworkStats{TxBytes: v(100)}, "prefix.network.tx_bytes", 100)
	checkNetworkStats(t, docker.NetworkStats{TxPackets: v(100)}, "prefix.network.tx_packets", 100)
	checkNetworkStats(t, docker.NetworkStats{TxErrors: v(100)}, "prefix.network.tx_errors", 100)
	checkNetworkStats(t, docker.NetworkStats{TxDropped: v(100)}, "prefix.network.tx_dropped", 100)
}

func checkNetworkStats(t *testing.T, stats docker.NetworkStats, name string, value int64) {
	checkTranslation(t, &docker.ContainerStats{Network: stats}, name, value)
}

func checkTranslation(t *testing.T, stats *docker.ContainerStats, name string, value int64) {
	stats.Timestamp = time.Date(2015, time.February, 23, 8, 39, 6, 0, time.UTC)
	metrics := Translate("prefix", stats)
	if len(metrics) != 1 {
		t.Fatalf("translation failed, should have only one metric: got %#v", metrics)
	}
	if metrics[0].Name != name {
		t.Errorf("translation failed: expected name '%s', got %s", name, metrics[0].Name)
	}
	if metrics[0].Value != value {
		t.Errorf("translation failed: expected value '%d', got %d", value, metrics[0].Value)
	}
	if metrics[0].Timestamp != stats.Timestamp.Unix() {
		t.Errorf("translation failed: expected timestamp '%d', got %d", stats.Timestamp.Unix(), metrics[0].Timestamp)
	}
}

func v(value int64) *int64 {
	return &value
}
