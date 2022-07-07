package collector

import (
	"encoding/json"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	urnMonitoring = "/scalaboom/widgets/MonitoringWidget"
)

type monitoringCollector struct {
	*xserver
}

type monitoringStatistics struct {
	Cpu   float64 `json:"cpu"`
	Ram   float64 `json:"ram"`
	Swap  float64 `json:"swap"`
	MaxLa float64 `json:"maxLa"`
	La    float64 `json:"la"`
}

var (
	monitoringCpuDesc = prometheus.NewDesc(
		"xserver_monitoring_cpu",
		"CPU Percentage",
		nil, nil,
	)
	monitoringRamDesc = prometheus.NewDesc(
		"xserver_monitoring_ram",
		"RAM percentage",
		nil, nil,
	)
	monitoringSwapDesc = prometheus.NewDesc(
		"xserver_monitoring_swap",
		"Swap usage",
		nil, nil,
	)
	monitoringMaxLaDesc = prometheus.NewDesc(
		"xserver_monitoring_max_la",
		"Max load average",
		nil, nil,
	)
	monitoringLaDesc = prometheus.NewDesc(
		"xserver_monitoring_la",
		"Load average",
		nil, nil,
	)
)

func (m monitoringCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(m, ch)
}

func (m monitoringCollector) Collect(ch chan<- prometheus.Metric) {
	ms := monitoringStatistics{}

	request, err := m.xserver.getJSON(urnMonitoring)
	if err != nil {
		m.logger.Fatalln(err)
	}

	err = json.Unmarshal(request, &ms)
	if err != nil {
		m.logger.Errorf("Failed unmarshal JSON due to err: %v", err)
	}

	ch <- prometheus.MustNewConstMetric(
		monitoringCpuDesc, prometheus.GaugeValue, ms.Cpu,
	)

	ch <- prometheus.MustNewConstMetric(
		monitoringRamDesc, prometheus.GaugeValue, ms.Ram,
	)

	ch <- prometheus.MustNewConstMetric(
		monitoringSwapDesc, prometheus.GaugeValue, ms.Swap,
	)

	ch <- prometheus.MustNewConstMetric(
		monitoringMaxLaDesc, prometheus.GaugeValue, ms.MaxLa,
	)

	ch <- prometheus.MustNewConstMetric(
		monitoringLaDesc, prometheus.GaugeValue, ms.La,
	)
}
