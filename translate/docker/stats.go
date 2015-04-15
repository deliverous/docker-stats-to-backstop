// vim: ts=2 nowrap
package docker

import (
	"time"
)

type ContainerStats struct {
	Timestamp time.Time    `json:"read"`
	Network   NetworkStats ``
	Cpu       CpuStats     `json:"cpu_stats"`
	Memory    MemoryStats  `json:"memory_stats"`
}

type NetworkStats struct {
	RxBytes   *uint64 `json:"rx_bytes"`
	RxPackets *uint64 `json:"rx_packets"`
	RxErrors  *uint64 `json:"rx_errors"`
	RxDropped *uint64 `json:"rx_dropped"`
	TxBytes   *uint64 `json:"tx_bytes"`
	TxPackets *uint64 `json:"tx_packets"`
	TxErrors  *uint64 `json:"tx_errors"`
	TxDropped *uint64 `json:"tx_dropped"`
}

type CpuStats struct {
	CpuUsage       CpuUsageStats `json:"cpu_usage"`
	SystemCpuUsage *uint64       `json:"system_cpu_usage"`
}

type CpuUsageStats struct {
	TotalUsage        *uint64 `json:"total_usage"`
	UsageInKernelmode *uint64 `json:"usage_in_kernelmode"`
	UsageInUsermode   *uint64 `json:"usage_in_usermode"`
}

type MemoryStats struct {
	Usage    *uint64            `json:"usage"`
	MaxUsage *uint64            `json:"max_usage"`
	Limit    *uint64            `json:"limit"`
	Stats    MemoryDetailsStats `json:"stats"`
}

type MemoryDetailsStats struct {
	TotalCache *uint64 `json:"total_cache"`
}
