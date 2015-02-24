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
	RxBytes   *int64 `json:"rx_bytes"`
	RxPackets *int64 `json:"rx_packets"`
	RxErrors  *int64 `json:"rx_errors"`
	RxDropped *int64 `json:"rx_dropped"`
	TxBytes   *int64 `json:"tx_bytes"`
	TxPackets *int64 `json:"tx_packets"`
	TxErrors  *int64 `json:"tx_errors"`
	TxDropped *int64 `json:"tx_dropped"`
}

type CpuStats struct {
	CpuUsage       CpuUsageStats `json:"cpu_usage"`
	SystemCpuUsage *int64        `json:"system_cpu_usage"`
}

type CpuUsageStats struct {
	TotalUsage        *int64 `json:"total_usage"`
	UsageInKernelmode *int64 `json:"usage_in_kernelmode"`
	UsageInUsermode   *int64 `json:"usage_in_usermode"`
}

type MemoryStats struct {
	Usage    *int64 `json:"usage"`
	MaxUsage *int64 `json:"max_usage"`
	Limit    *int64 `json:"limit"`
}
