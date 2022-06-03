package collector

import (
	"log"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type IfacesWidget struct {
	Today `json:"today"`
}

type Today struct {
	Ping           string `json:"ping"`
	VpnConnections int    `json:"vpn_connections"`
	ChannelUsage `json:"channelUsage"`
}

type ChannelUsage struct {
	Rp float32 `json:"rp"`
	Tp float32 `json:"tp"`
	Rx float32 `json:"rx"`
	Tx float32 `json:"tx"`
}

var (
	netstatPingDesc = prometheus.NewDesc(
		"xserver_netstat_ping",
		"Ping to ya.ru",
		nil, nil,
	)
	netstatVpnConnectionsDesc = prometheus.NewDesc(
		"xserver_netstat_vpn_connections",
		"Тumber of active vpn connections",
		nil, nil,
	)
	netstatChannelUsageRpDesc = prometheus.NewDesc(
		"xserver_netstat_channel_usage_rp",
		"Received packets per second",
		nil, nil,
	)
	netstatChannelUsageTpDesc = prometheus.NewDesc(
		"xserver_netstat_channel_usage_tp",
		"Transmit packets per second",
		nil, nil,
	)
	netstatChannelUsageRxDesc = prometheus.NewDesc(
		"xserver_netstat_channel_usage_rx",
		"Received bytes per second",
		nil, nil,
	)
	netstatChannelUsageTxDesc = prometheus.NewDesc(
		"xserver_netstat_channel_usage_tx",
		"Transmit bytes per second",
		nil, nil,
	)
)

func (i IfacesWidget) Collect(ch chan<- prometheus.Metric) {
	ping, err := strconv.ParseFloat(i.Ping, 64)
	if err != nil {
        log.Fatal(err)
    }

	ch <- prometheus.MustNewConstMetric(
		netstatPingDesc, prometheus.GaugeValue, ping,
	)

	ch <- prometheus.MustNewConstMetric(
		netstatVpnConnectionsDesc, prometheus.GaugeValue, float64(i.VpnConnections),
	)

	ch <- prometheus.MustNewConstMetric(
		netstatChannelUsageRpDesc, prometheus.GaugeValue, float64(i.Rp),
	)

	ch <- prometheus.MustNewConstMetric(
		netstatChannelUsageTpDesc, prometheus.GaugeValue, float64(i.Tp),
	)

	ch <- prometheus.MustNewConstMetric(
		netstatChannelUsageRxDesc, prometheus.GaugeValue, float64(i.Rx),
	)

	ch <- prometheus.MustNewConstMetric(
		netstatChannelUsageTxDesc, prometheus.GaugeValue, float64(i.Tx),
	)
}