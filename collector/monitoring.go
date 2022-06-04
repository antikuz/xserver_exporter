package collector

import "github.com/prometheus/client_golang/prometheus"

type monitoringWidget struct {
	Cpu   float64 `json:"cpu"`
	Ram   float64 `json:"ram"`
	Swap  float64 `json:"swap"`
	MaxLa float64 `json:"maxLa"`
	La    float64 `json:"la"`
}

var (
	MonitoringWidget monitoringWidget
	
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

func (m monitoringWidget) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		monitoringCpuDesc, prometheus.GaugeValue, m.Cpu,
	)

	ch <- prometheus.MustNewConstMetric(
		monitoringRamDesc, prometheus.GaugeValue, m.Ram,
	)

	ch <- prometheus.MustNewConstMetric(
		monitoringSwapDesc, prometheus.GaugeValue, m.Swap,
	)

	ch <- prometheus.MustNewConstMetric(
		monitoringMaxLaDesc, prometheus.GaugeValue, m.MaxLa,
	)

	ch <- prometheus.MustNewConstMetric(
		monitoringLaDesc, prometheus.GaugeValue, m.La,
	)
}