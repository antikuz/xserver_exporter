package collector

import (
	"encoding/json"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	urnIfaces      = "/scalaboom/widgets/IfacesWidget"
	urnNetstatStat = "/scalaboom/widgets/StatWidget"
)

type netstatCollector struct {
	*xserver
}

// type netstat struct {
// 	IfacesWidget
// 	StatWidget
// }

type IfacesWidget struct {
	Today `json:"today"`
}

type Today struct {
	Ping           string `json:"ping"`
	VpnConnections int    `json:"vpn_connections"`
	ChannelUsage   `json:"channelUsage"`
}

type ChannelUsage struct {
	Rp float64 `json:"rp"`
	Tp float64 `json:"tp"`
	Rx float64 `json:"rx"`
	Tx float64 `json:"tx"`
}

type StatWidget struct {
	Stat `json:"stat"`
}

type Stat struct {
	RxMonth float64 `json:"rxMonth"`
	TxMonth float64 `json:"txMonth"`
	RxWeek  float64 `json:"rxWeek"`
	TxWeek  float64 `json:"txWeek"`
	RxDay   float64 `json:"rxDay"`
	TxDay   float64 `json:"txDay"`
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
	netstatStatRxMonthDesc = prometheus.NewDesc(
		"xserver_netstat_stat_rx_month",
		"Received bytes per month",
		nil, nil,
	)
	netstatStatTxMonthDesc = prometheus.NewDesc(
		"xserver_netstat_stat_tx_month",
		"Transmit bytes per month",
		nil, nil,
	)
	netstatStatRxWeekDesc = prometheus.NewDesc(
		"xserver_netstat_stat_rx_week",
		"Received bytes per week",
		nil, nil,
	)
	netstatStatTxWeekDesc = prometheus.NewDesc(
		"xserver_netstat_stat_tx_week",
		"Transmit bytes per week",
		nil, nil,
	)
	netstatStatRxDayDesc = prometheus.NewDesc(
		"xserver_netstat_stat_rx_day",
		"Received bytes per day",
		nil, nil,
	)
	netstatStatTxDayDesc = prometheus.NewDesc(
		"xserver_netstat_stat_tx_day",
		"Transmit bytes per day",
		nil, nil,
	)
)

func (n netstatCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(n, ch)
}

func (n netstatCollector) Collect(ch chan<- prometheus.Metric) {
	iw := IfacesWidget{}

	request, err := n.xserver.getJSON(urnIfaces)
	if err != nil {
		n.logger.Fatalln(err)
	}
	json.Unmarshal(request, &iw)

	sw := StatWidget{}

	request, err = n.xserver.getJSON(urnNetstatStat)
	if err != nil {
		n.logger.Fatalln(err)
	}
	json.Unmarshal(request, &sw)

	ping, err := strconv.ParseFloat(iw.Ping, 64)
	if err != nil {
		n.logger.Fatalln(err)
	}

	ch <- prometheus.MustNewConstMetric(
		netstatPingDesc, prometheus.GaugeValue, ping,
	)

	ch <- prometheus.MustNewConstMetric(
		netstatVpnConnectionsDesc, prometheus.GaugeValue, float64(iw.VpnConnections),
	)

	ch <- prometheus.MustNewConstMetric(
		netstatChannelUsageRpDesc, prometheus.GaugeValue, iw.Rp,
	)

	ch <- prometheus.MustNewConstMetric(
		netstatChannelUsageTpDesc, prometheus.GaugeValue, iw.Tp,
	)

	ch <- prometheus.MustNewConstMetric(
		netstatChannelUsageRxDesc, prometheus.GaugeValue, iw.Rx,
	)

	ch <- prometheus.MustNewConstMetric(
		netstatChannelUsageTxDesc, prometheus.GaugeValue, iw.Tx,
	)

	ch <- prometheus.MustNewConstMetric(
		netstatStatRxMonthDesc, prometheus.GaugeValue, sw.RxMonth,
	)

	ch <- prometheus.MustNewConstMetric(
		netstatStatTxMonthDesc, prometheus.GaugeValue, sw.TxMonth,
	)

	ch <- prometheus.MustNewConstMetric(
		netstatStatRxWeekDesc, prometheus.GaugeValue, sw.RxWeek,
	)

	ch <- prometheus.MustNewConstMetric(
		netstatStatTxWeekDesc, prometheus.GaugeValue, sw.TxWeek,
	)

	ch <- prometheus.MustNewConstMetric(
		netstatStatRxDayDesc, prometheus.GaugeValue, sw.RxDay,
	)

	ch <- prometheus.MustNewConstMetric(
		netstatStatTxDayDesc, prometheus.GaugeValue, sw.TxDay,
	)
}
