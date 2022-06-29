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
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.GetConfig()
	
	logger := logging.GetLogger()
	loglevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Logger.SetLevel(loglevel)
	
	registry := prometheus.NewPedanticRegistry()
	collector.NewXserverManager(
		logger,
		cfg.Url,
		cfg.Login,
		cfg.Passwd,
		cfg.Insecure,
		registry,
	)

	registry.MustRegister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(),
	)

	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
