package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/percona/rds_exporter/basic"
	"github.com/percona/rds_exporter/client"
	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/enhanced"
	"github.com/percona/rds_exporter/sessions"
)

var (
	listenAddress       = kingpin.Flag("web.listen-address", "Address on which to expose metrics and web interface.").Default(":9042").String()
	basicMetricsPath    = kingpin.Flag("web.basic-telemetry-path", "Path under which to expose exporter's basic metrics.").Default("/basic").String()
	enhancedMetricsPath = kingpin.Flag("web.enhanced-telemetry-path", "Path under which to expose exporter's enhanced metrics.").Default("/enhanced").String()
	configFile          = kingpin.Flag("config.file", "Path to configuration file.").Default("config.yml").String()
)

func main() {
	log.AddFlags(kingpin.CommandLine)
	log.Infoln("Starting RDS exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	kingpin.Parse()

	config, err := config.Load(*configFile)
	if err != nil {
		log.Fatalf("Can't read configuration file: %s", err)
	}
	client := client.New(logrus.WithField("component", "client"))
	sessions, err := sessions.New(config, client, logrus.WithField("component", "sessions"))
	if err != nil {
		log.Fatalf("Can't create sessions: %s", err)
	}

	// Basic Metrics
	{
		// Create new Exporter with provided settings and session pool.
		exporter := basic.New(config, sessions)
		registry := prometheus.NewRegistry()
		registry.MustRegister(exporter)
		handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		// Expose the exporter's own metrics on given path
		http.Handle(*basicMetricsPath, handler)
	}

	// Enhanced Metrics
	{
		// Create new Exporter with provided settings and session pool.
		exporter := enhanced.New(config, sessions)
		registry := prometheus.NewRegistry()
		registry.MustRegister(exporter)
		handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		// Expose the exporter's own metrics on given path
		http.Handle(*enhancedMetricsPath, handler)
	}

	// Inform user we are ready.
	log.Infoln("RDS exporter started on", *listenAddress)

	// Start serving for clients
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
