package collector

import (
	"encoding/json"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	urnUsers = "/scalaboom/widgets/UsersWidget"
)

type usersCollector struct {
	*xserver
}

type usersWidget struct {
	Active  int `json:"active"`
	Blocked int `json:"blocked"`
	Enabled int `json:"enabled"`
	Users   int `json:"users"`
	Vpn     int `json:"vpn"`
}

var (
	usersActiveDesc = prometheus.NewDesc(
		"xserver_users_active",
		"Number of users online.",
		nil, nil,
	)
	usersBlockedDesc = prometheus.NewDesc(
		"xserver_users_blocked",
		"Number of users blocked.",
		nil, nil,
	)
	usersEnabledDesc = prometheus.NewDesc(
		"xserver_users_enabled",
		"Number of users enabled.",
		nil, nil,
	)
	usersTotalDesc = prometheus.NewDesc(
		"xserver_users_total",
		"Number of users total.",
		nil, nil,
	)

	usersVpnDesc = prometheus.NewDesc(
		"xserver_users_vpn",
		"Number of users connected by vpn.",
		nil, nil,
	)
)

func (u usersCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(u, ch)
}

func (u usersCollector) Collect(ch chan<- prometheus.Metric) {
	uw := usersWidget{}

	request, err := u.xserver.getJSON(urnUsers)
	if err != nil {
		u.logger.Fatalln(err)
	}

	err = json.Unmarshal(request, &uw)
	if err != nil {
		u.logger.Errorf("UsersCollect failed unmarshal JSON due to err: %v", err)
	}

	ch <- prometheus.MustNewConstMetric(
		usersActiveDesc, prometheus.GaugeValue, float64(uw.Active),
	)

	ch <- prometheus.MustNewConstMetric(
		usersBlockedDesc, prometheus.GaugeValue, float64(uw.Blocked),
	)

	ch <- prometheus.MustNewConstMetric(
		usersEnabledDesc, prometheus.GaugeValue, float64(uw.Enabled),
	)

	ch <- prometheus.MustNewConstMetric(
		usersTotalDesc, prometheus.GaugeValue, float64(uw.Users),
	)

	ch <- prometheus.MustNewConstMetric(
		usersVpnDesc, prometheus.GaugeValue, float64(uw.Vpn),
	)
}
