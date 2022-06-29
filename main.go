package main

import (
	"log"
	"net/http"

	"github.com/antikuz/xserver_exporter/collector"
	"github.com/antikuz/xserver_exporter/internal/config"
	"github.com/antikuz/xserver_exporter/pkg/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()

	registry := prometheus.NewPedanticRegistry()
	collector.NewXserverManager(
		logger,
		cfg.Url,
		cfg.Login,
		cfg.Passwd,
		cfg.InsecureSkip,
		registry,
	)

	registry.MustRegister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(),
	)

	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
