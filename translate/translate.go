// vim: ts=2 nowrap
package translate

import (
	"github.com/deliverous/collectd-docker/translate/backstop"
	"github.com/deliverous/collectd-docker/translate/docker"
	"time"
)

func Translate(prefix string, stats *docker.ContainerStats) []backstop.Metric {
	c := collector{Prefix: prefix, Time: time.Now(), Metrics: []backstop.Metric{}}
	c.add("cpu.system", stats.Cpu.SystemCpuUsage)
	c.add("cpu.total", stats.Cpu.CpuUsage.TotalUsage)
	c.add("cpu.kernel", stats.Cpu.CpuUsage.UsageInKernelmode)
	c.add("cpu.user", stats.Cpu.CpuUsage.UsageInUsermode)
	c.add("memory.usage", stats.Memory.Usage)
	c.add("memory.max_usage", stats.Memory.MaxUsage)
	c.add("memory.limit", stats.Memory.Limit)
	c.add("network.rx_bytes", stats.Network.RxBytes)
	c.add("network.rx_packets", stats.Network.RxPackets)
	c.add("network.rx_errors", stats.Network.RxErrors)
	c.add("network.rx_dropped", stats.Network.RxDropped)
	c.add("network.tx_bytes", stats.Network.TxBytes)
	c.add("network.tx_packets", stats.Network.TxPackets)
	c.add("network.tx_errors", stats.Network.TxErrors)
	c.add("network.tx_dropped", stats.Network.TxDropped)

	return c.Metrics
}

type collector struct {
	Prefix  string
	Time    time.Time
	Metrics []backstop.Metric
}

func (c *collector) add(name string, value *int64) {
	if value != nil {
		c.Metrics = append(c.Metrics, backstop.Metric{
			Name:  c.Prefix + "." + name,
			Value: *value,
			Time:  c.Time,
		})
	}
}
