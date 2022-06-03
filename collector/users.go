package collector

import "github.com/prometheus/client_golang/prometheus"

type UsersWidget struct {
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
		[]string{}, nil,
	)
	usersBlockedDesc = prometheus.NewDesc(
		"xserver_users_blocked",
		"Number of users blocked.",
		[]string{}, nil,
	)
	usersEnabledDesc = prometheus.NewDesc(
		"xserver_users_enabled",
		"Number of users enabled.",
		[]string{}, nil,
	)
	usersTotalDesc = prometheus.NewDesc(
		"xserver_users_total",
		"Number of users total.",
		[]string{}, nil,
	)

	usersVpnDesc = prometheus.NewDesc(
		"xserver_users_total",
		"Number of users connected by vpn.",
		[]string{}, nil,
	)
)

func (u UsersWidget) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		usersActiveDesc, prometheus.GaugeValue, float64(u.Active),
	)

	ch <- prometheus.MustNewConstMetric(
		usersBlockedDesc, prometheus.GaugeValue, float64(u.Blocked),
	)

	ch <- prometheus.MustNewConstMetric(
		usersEnabledDesc, prometheus.GaugeValue, float64(u.Enabled),
	)

	ch <- prometheus.MustNewConstMetric(
		usersTotalDesc, prometheus.GaugeValue, float64(u.Users),
	)

	ch <- prometheus.MustNewConstMetric(
		usersVpnDesc, prometheus.GaugeValue, float64(u.Vpn),
	)
}