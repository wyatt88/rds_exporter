package enhanced

import (
	"testing"

	"github.com/percona/exporter_shared/helpers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/sessions"
)

func getExporter(t *testing.T) *Exporter {
	t.Helper()

	config, err := config.Load("../config.yml")
	require.NoError(t, err)
	sessions, err := sessions.Load(config.Instances)
	require.NoError(t, err)
	return New(config, sessions)
}

func TestCollector_Collect(t *testing.T) {
	t.Parallel()

	c := getExporter(t)
	ch := make(chan prometheus.Metric)
	go func() {
		c.Collect(ch)
		close(ch)
	}()

	var metrics []helpers.Metric
	for m := range ch {
		metrics = append(metrics, *helpers.ReadMetric(m))
	}

	assert.Len(t, metrics, 3, "%+v", metrics)
}

func TestCollector_Describe(t *testing.T) {
	t.Parallel()

	c := getExporter(t)
	ch := make(chan *prometheus.Desc)
	go func() {
		c.Describe(ch)
		close(ch)
	}()

	var descs []*prometheus.Desc
	for d := range ch {
		descs = append(descs, d)
	}

	assert.Len(t, descs, 81, "%+v", descs)
}
