package enhanced2

import (
	"testing"
	"time"

	"github.com/percona/exporter_shared/helpers"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	m, err := parseOSMetrics(dataMySQL57)
	require.NoError(t, err)
	timestamp := time.Date(2018, 1, 8, 14, 35, 21, 0, time.UTC)
	assert.Equal(t, timestamp, m.Timestamp)

	metrics := m.originalMetrics("prefix", "aws-region")
	assert.Len(t, metrics, 270)

	for i, m := range metrics {
		t.Logf("%d: %#v", i, helpers.ReadMetric(m))
	}

	for i, expected := range map[int]*helpers.Metric{
		6: &helpers.Metric{
			Name:   "prefix_cpuUtilization_total",
			Help:   "The total percentage of the CPU in use. This value includes the nice value.",
			Labels: prometheus.Labels{"instance": "rds-mysql57", "region": "aws-region"},
			Type:   dto.MetricType_GAUGE,
			Value:  99.99,
		},
		13: &helpers.Metric{
			Name:   "prefix_diskIO_readKb",
			Help:   "The total number of kilobytes read.",
			Labels: prometheus.Labels{"instance": "rds-mysql57", "region": "aws-region", "device": "rdsdev"},
			Type:   dto.MetricType_GAUGE,
			Value:  339896,
		},
		26: &helpers.Metric{
			Name:   "prefix_fileSys_usedFiles",
			Help:   "The number of files in the file system.",
			Labels: prometheus.Labels{"instance": "rds-mysql57", "region": "aws-region", "name": "rdsfilesys", "mount_point": "/rdsdbdata"},
			Type:   dto.MetricType_GAUGE,
			Value:  731,
		},
		29: &helpers.Metric{
			Name:   "prefix_loadAverageMinute_five",
			Help:   "The number of processes requesting CPU time over the last 5 minutes.",
			Labels: prometheus.Labels{"instance": "rds-mysql57", "region": "aws-region"},
			Type:   dto.MetricType_GAUGE,
			Value:  14.08,
		},
		41: &helpers.Metric{
			Name:   "prefix_memory_inactive",
			Help:   "The amount of least-frequently used memory pages, in kilobytes.",
			Labels: prometheus.Labels{"instance": "rds-mysql57", "region": "aws-region"},
			Type:   dto.MetricType_GAUGE,
			Value:  511232,
		},
		48: &helpers.Metric{
			Name:   "prefix_network_tx",
			Help:   "The number of bytes uploaded per second.",
			Labels: prometheus.Labels{"instance": "rds-mysql57", "region": "aws-region", "interface": "eth0"},
			Type:   dto.MetricType_GAUGE,
			Value:  801373.55,
		},
		250: &helpers.Metric{
			Name:   "prefix_processList_memoryUsedPc",
			Help:   "The amount of memory used by the process, in kilobytes.",
			Labels: prometheus.Labels{"region": "aws-region", "tgid": "4516", "id": "4551", "instance": "rds-mysql57", "name": "mysqld", "parentID": "1"},
			Type:   dto.MetricType_GAUGE,
			Value:  62.51,
		},
		262: &helpers.Metric{
			Name:   "prefix_swap_free",
			Help:   "The total amount of swap memory free, in kilobytes.",
			Labels: prometheus.Labels{"instance": "rds-mysql57", "region": "aws-region"},
			Type:   dto.MetricType_GAUGE,
			Value:  3803212,
		},
		268: &helpers.Metric{
			Name:   "prefix_tasks_total",
			Help:   "The total number of tasks.",
			Labels: prometheus.Labels{"instance": "rds-mysql57", "region": "aws-region"},
			Type:   dto.MetricType_GAUGE,
			Value:  313,
		},
	} {
		assert.Equal(t, expected, helpers.ReadMetric(metrics[i]))
	}
}
