package main

import (
	"log"
	"net/http"

	"github.com/antikuz/xserver_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	url := "http://localhost:8000"
	login := "login"
	passwd := "passwd"
	insecureSkip := true

	reg := prometheus.NewPedanticRegistry()
	collector.NewXserverManager(
		url,
		login,
		passwd,
		insecureSkip,
		reg)

	reg.MustRegister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(),
	)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
