package enhanced

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"

	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/latency"
	"github.com/percona/rds_exporter/sessions"
)

//go:generate go run generate/main.go

const (
	defaultNamespace = "rdsosmetrics"
	logGroupName     = "RDSOSMetrics"
)

var (
	ScrapeTimeDesc = prometheus.NewDesc(
		"rds_exporter_scrape_duration_seconds",
		"Time this RDS scrape took, in seconds.",
		[]string{},
		nil,
	)
)

type Metric struct {
	Name string
	Desc *prometheus.Desc
}

type Exporter struct {
	Config   *config.Config
	Sessions *sessions.Sessions

	// Metrics
	Metrics           map[string]Metric
	ErroneousRequests prometheus.Counter
	TotalRequests     prometheus.Counter

	l *logrus.Entry
}

func New(config *config.Config, sessions *sessions.Sessions) *Exporter {
	return &Exporter{
		Config:   config,
		Sessions: sessions,
		Metrics:  Metrics,

		ErroneousRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "rds_exporter_erroneous_requests",
			Help: "The number of erroneous API request made to CloudWatch.",
		}),
		TotalRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "rds_exporter_requests_total",
			Help: "API requests made to CloudWatch",
		}),

		l: logrus.WithField("component", "enhanced"),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	// RDS metrics
	for _, m := range e.Metrics {
		ch <- m.Desc
	}

	// Latency metric
	ch <- latency.Desc

	// Scrape time
	ch <- ScrapeTimeDesc

	// Global number of requests, and global number of failed requests
	ch <- e.TotalRequests.Desc()
	ch <- e.ErroneousRequests.Desc()
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	now := time.Now()
	e.collect(ch)

	// Collect scrape time
	ch <- prometheus.MustNewConstMetric(ScrapeTimeDesc, prometheus.GaugeValue, float64(time.Since(now).Seconds()))

	// Collect global number of requests, and global number of failed requests
	ch <- e.TotalRequests
	ch <- e.ErroneousRequests
}

func (e *Exporter) collect(ch chan<- prometheus.Metric) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	instances := e.Config.Instances
	wg.Add(len(instances))
	for _, instance := range instances {
		go func(instance *config.Instance) {
			defer wg.Done()

			err := e.collectInstance(ch, instance)
			if err != nil {
				e.l.WithField("region", instance.Region).WithField("instance", instance.Instance).Error(err)
			}
		}(instance)
	}
}

func (e *Exporter) collectInstance(ch chan<- prometheus.Metric, instance *config.Instance) error {
	// Latency metric
	l := &latency.Latency{}

	// Collect all values
	err := e.collectValues(ch, instance, l)
	if err != nil {
		return err
	}

	// Collect latency metric
	if !l.IsZero() {
		ch <- prometheus.MustNewConstMetric(latency.Desc, prometheus.GaugeValue, float64(l.Duration().Seconds()), instanceLabels(instance)...)
	}
	return nil
}

func (e *Exporter) collectValues(ch chan<- prometheus.Metric, instance *config.Instance, l *latency.Latency) error {
	sess := e.Sessions.Get(instance)
	svc := cloudwatchlogs.New(sess)

	FilterLogEventsOutput, err := svc.FilterLogEvents(&cloudwatchlogs.FilterLogEventsInput{
		Limit:         aws.Int64(1),
		LogGroupName:  aws.String(logGroupName),
		FilterPattern: aws.String(fmt.Sprintf(`{ $.instanceID = "%s" }`, instance.Instance)),
	})
	if err != nil {
		return fmt.Errorf("unable to get logs: %s", err)
	}

	if len(FilterLogEventsOutput.Events) == 0 {
		return errors.New("no log events")
	}

	var message interface{}
	err = json.Unmarshal([]byte(*FilterLogEventsOutput.Events[0].Message), &message)
	if err != nil {
		return err
	}

	values, ok := message.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unsupported message type: %T", message)
	}

	if len(values) == 0 {
		return nil
	}

	wg := &sync.WaitGroup{}
	defer wg.Wait()

	wg.Add(len(values))
	for key, value := range values {
		go func(key string, value interface{}) {
			defer wg.Done()

			err := e.collectValue(ch, instance, key, value, l)
			if err != nil {
				e.l.WithField("region", instance.Region).WithField("instance", instance.Instance).WithField("key", key).Error(err)
			}
		}(key, value)
	}

	return nil
}

func (e *Exporter) collectValue(ch chan<- prometheus.Metric, instance *config.Instance, key string, value interface{}, l *latency.Latency) error {
	switch v := value.(type) {
	case float64:
		e.sendMetric(ch, instance, defaultNamespace, "General", key, v)
	case map[string]interface{}:
		e.collectMapValue(ch, instance, key, v, l)
	case []interface{}:
		for i, u := range v {
			extraLabels := []string{
				strconv.Itoa(i),
			}
			metrics, ok := u.(map[string]interface{})
			if !ok {
				return fmt.Errorf("%s: wrong value type for metrics: %T", key, metrics)
			}
			e.collectMapValue(ch, instance, key, metrics, l, extraLabels...)
		}
	case string:
		switch key {
		case
			"engine",
			"instanceID",
			"instanceResourceID",
			"uptime":
			// skipping those values
		case "timestamp":
			t, err := time.Parse(time.RFC3339, v)
			if err != nil {
				return err
			}
			l.TakeOldest(t)
		default:
			return fmt.Errorf("unsupported key '%s' when collecting value", key)
		}
	default:
		return fmt.Errorf("unsupported value type '%T' for key '%s' when collecting value", value, key)
	}

	return nil
}

func (e *Exporter) collectMapValue(ch chan<- prometheus.Metric, instance *config.Instance, key string, value map[string]interface{}, l *latency.Latency, labels ...string) error {
	for metricName, v := range value {
		metricValue, ok := v.(float64)
		if !ok {
			return fmt.Errorf("%s: wrong value type for metric %s: %T", key, metricName, value)
		}

		namespace, subsystem, name, labels := MapToNode(key, metricName, labels...)
		e.sendMetric(ch, instance, namespace, subsystem, name, metricValue, labels...)
	}

	return nil
}

func (e *Exporter) sendMetric(ch chan<- prometheus.Metric, instance *config.Instance, namespace, subsystem, name string, value float64, extraLabels ...string) {
	FQName := prometheus.BuildFQName(namespace, subsystem, name)
	metric, ok := e.Metrics[FQName]
	if !ok {
		e.l.WithField("region", instance.Region).WithField("instance", instance.Instance).Errorf("unknown metric %s", FQName)
		return
	}

	labels := instanceLabels(instance)
	if len(extraLabels) > 0 {
		labels = append(labels, extraLabels...)
	}

	ch <- prometheus.MustNewConstMetric(
		metric.Desc,
		prometheus.GaugeValue,
		value,
		labels...,
	)
}

func instanceLabels(instance *config.Instance) []string {
	return []string{
		instance.Instance,
		instance.Region,
	}
}
