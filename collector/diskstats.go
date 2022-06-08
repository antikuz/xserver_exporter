package collector

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	urnMainHDD = "/scalaboom/harddisks/MainSection1"
)

type harddisksCollector struct {
	*xserver
}

type harddisksStatistics struct {
	Name   string `json:"name"`
	Used   string `json:"used"`
	Avail  string `json:"avail"`
	Size   string `json:"size"`
}

var (
	harddisksUsedDesc = prometheus.NewDesc(
		"xserver_harddisks_used",
		"CPU Percentage",
		[]string{"name"}, nil,
	)
	harddisksAvailDesc = prometheus.NewDesc(
		"xserver_harddisks_avail",
		"CPU Percentage",
		[]string{"name"}, nil,
	)
	harddisksSizeDesc = prometheus.NewDesc(
		"xserver_harddisks_size",
		"CPU Percentage",
		[]string{"name"}, nil,
	)
)

func diskSizeToFloat(size string) float64 {
	number := strings.TrimRight(size, " ГG")
	numberFloat, err := strconv.ParseFloat(number, 64)
	if err != nil {
		log.Fatal(err)
	}

	return numberFloat
}

func (h harddisksCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(h, ch)
}

func (h harddisksCollector) Collect(ch chan<- prometheus.Metric) {
	hs := harddisksStatistics{}

	request, err := h.xserver.getJSON(urnMainHDD)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(request, &hs)

	ch <- prometheus.MustNewConstMetric(
		harddisksUsedDesc, prometheus.GaugeValue, diskSizeToFloat(hs.Used), hs.Name,
	)

	ch <- prometheus.MustNewConstMetric(
		harddisksAvailDesc, prometheus.GaugeValue, diskSizeToFloat(hs.Avail), hs.Name,
	)

	ch <- prometheus.MustNewConstMetric(
		harddisksSizeDesc, prometheus.GaugeValue, diskSizeToFloat(hs.Size), hs.Name,
	)

}
