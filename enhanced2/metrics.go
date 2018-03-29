package enhanced2

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// osMetrics represents available Enhanced Monitoring OS metrics from CloudWatch Logs.
//
// See "Metrics for Amazon Aurora, MariaDB, MySQL, Oracle, and PostgreSQL DB instances" from
// https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.CloudWatchLogs
type osMetrics struct {
	Engine             string    `json:"engine"             help:"The database engine for the DB instance."`
	InstanceID         string    `json:"instanceID"         help:"The DB instance identifier."`
	InstanceResourceID string    `json:"instanceResourceID" help:"A region-unique, immutable identifier for the DB instance, also used as the log stream identifier."`
	NumVCPUs           int       `json:"numVCPUs"           help:"The number of virtual CPUs for the DB instance."`
	Timestamp          time.Time `json:"timestamp"          help:"The time at which the metrics were taken."`
	Uptime             string    `json:"uptime"             help:"The amount of time that the DB instance has been active."`
	Version            float64   `json:"version"            help:"The version of the OS metrics' stream JSON format."`

	CPUUtilization struct {
		Guest  float64 `json:"guest"  help:"The percentage of CPU in use by guest programs."`
		Idle   float64 `json:"idle"   help:"The percentage of CPU that is idle."`
		Irq    float64 `json:"irq"    help:"The percentage of CPU in use by software interrupts."`
		Nice   float64 `json:"nice"   help:"The percentage of CPU in use by programs running at lowest priority."`
		Steal  float64 `json:"steal"  help:"The percentage of CPU in use by other virtual machines."`
		System float64 `json:"system" help:"The percentage of CPU in use by the kernel."`
		Total  float64 `json:"total"  help:"The total percentage of the CPU in use. This value includes the nice value."`
		User   float64 `json:"user"   help:"The percentage of CPU in use by user programs."`
		Wait   float64 `json:"wait"   help:"The percentage of CPU unused while waiting for I/O access."`
	} `json:"cpuUtilization"`

	// absent for Aurora
	DiskIO []struct {
		AvgQueueLen float64 `json:"avgQueueLen" help:"The number of requests waiting in the I/O device's queue."`
		AvgReqSz    float64 `json:"avgReqSz"    help:"The average request size, in kilobytes."`
		Await       float64 `json:"await"       help:"The number of milliseconds required to respond to requests, including queue time and service time."`
		Device      string  `json:"device"      help:"The identifier of the disk device in use."`
		ReadIOsPS   float64 `json:"readIOsPS"   help:"The number of read operations per second."`
		ReadKb      int     `json:"readKb"      help:"The total number of kilobytes read."`
		ReadKbPS    float64 `json:"readKbPS"    help:"The number of kilobytes read per second."`
		RrqmPS      float64 `json:"rrqmPS"      help:"The number of merged read requests queued per second."`
		TPS         float64 `json:"tps"         help:"The number of I/O transactions per second."`
		Util        float64 `json:"util"        help:"The percentage of CPU time during which requests were issued."`
		WriteIOsPS  float64 `json:"writeIOsPS"  help:"The number of write operations per second."`
		WriteKb     int     `json:"writeKb"     help:"The total number of kilobytes written."`
		WriteKbPS   float64 `json:"writeKbPS"   help:"The number of kilobytes written per second."`
		WrqmPS      float64 `json:"wrqmPS"      help:"The number of merged write requests queued per second."`
	} `json:"diskIO"`

	FileSys []struct {
		MaxFiles        int     `json:"maxFiles"        help:"The maximum number of files that can be created for the file system."`
		MountPoint      string  `json:"mountPoint"      help:"The path to the file system."`
		Name            string  `json:"name"            help:"The name of the file system."`
		Total           int     `json:"total"           help:"The total number of disk space available for the file system, in kilobytes."`
		Used            int     `json:"used"            help:"The amount of disk space used by files in the file system, in kilobytes."`
		UsedFilePercent float64 `json:"usedFilePercent" help:"The percentage of available files in use."`
		UsedFiles       int     `json:"usedFiles"       help:"The number of files in the file system."`
		UsedPercent     float64 `json:"usedPercent"     help:"The percentage of the file-system disk space in use."`
	} `json:"fileSys"`

	LoadAverageMinute struct {
		Fifteen float64 `json:"fifteen" help:"The number of processes requesting CPU time over the last 15 minutes."`
		Five    float64 `json:"five"    help:"The number of processes requesting CPU time over the last 5 minutes."`
		One     float64 `json:"one"     help:"The number of processes requesting CPU time over the last minute.0"`
	} `json:"loadAverageMinute"`

	Memory struct {
		Active         int `json:"active"         help:"The amount of assigned memory, in kilobytes."`
		Buffers        int `json:"buffers"        help:"The amount of memory used for buffering I/O requests prior to writing to the storage device, in kilobytes."`
		Cached         int `json:"cached"         help:"The amount of memory used for caching file system–based I/O."`
		Dirty          int `json:"dirty"          help:"The amount of memory pages in RAM that have been modified but not written to their related data block in storage, in kilobytes."`
		Free           int `json:"free"           help:"The amount of unassigned memory, in kilobytes."`
		HugePagesFree  int `json:"hugePagesFree"  help:"The number of free huge pages. Huge pages are a feature of the Linux kernel."`
		HugePagesRsvd  int `json:"hugePagesRsvd"  help:"The number of committed huge pages."`
		HugePagesSize  int `json:"hugePagesSize"  help:"The size for each huge pages unit, in kilobytes."`
		HugePagesSurp  int `json:"hugePagesSurp"  help:"The number of available surplus huge pages over the total."`
		HugePagesTotal int `json:"hugePagesTotal" help:"The total number of huge pages for the system."`
		Inactive       int `json:"inactive"       help:"The amount of least-frequently used memory pages, in kilobytes."`
		Mapped         int `json:"mapped"         help:"The total amount of file-system contents that is memory mapped inside a process address space, in kilobytes."`
		PageTables     int `json:"pageTables"     help:"The amount of memory used by page tables, in kilobytes."`
		Slab           int `json:"slab"           help:"The amount of reusable kernel data structures, in kilobytes."`
		Total          int `json:"total"          help:"The total amount of memory, in kilobytes."`
		Writeback      int `json:"writeback"      help:"The amount of dirty pages in RAM that are still being written to the backing storage, in kilobytes."`
	} `json:"memory"`

	Network []struct {
		Interface string  `json:"interface" help:"The identifier for the network interface being used for the DB instance."`
		Rx        float64 `json:"rx"        help:"The number of bytes received per second."`
		Tx        float64 `json:"tx"        help:"The number of bytes uploaded per second."`
	} `json:"network"`

	ProcessList []struct {
		CPUUsedPC    float64 `json:"cpuUsedPc"    help:"The percentage of CPU used by the process."`
		ID           int     `json:"id"           help:"The identifier of the process."`
		MemoryUsedPC float64 `json:"memoryUsedPc" help:"The amount of memory used by the process, in kilobytes."`
		Name         string  `json:"name"         help:"The name of the process."`
		ParentID     int     `json:"parentID"     help:"The process identifier for the parent process of the process."`
		RSS          int     `json:"rss"          help:"The amount of RAM allocated to the process, in kilobytes."`
		TGID         int     `json:"tgid"         help:"The thread group identifier, which is a number representing the process ID to which a thread belongs. This identifier is used to group threads from the same process."`
		VSS          int     `json:"vss"          help:"The amount of virtual memory allocated to the process, in kilobytes."`
	} `json:"processList"`

	Swap struct {
		Cached int `json:"cached" help:"The amount of swap memory, in kilobytes, used as cache memory."`
		Free   int `json:"free"   help:"The total amount of swap memory free, in kilobytes."`
		Total  int `json:"total"  help:"The total amount of swap memory available, in kilobytes."`
	} `json:"swap"`

	Tasks struct {
		Blocked  int `json:"blocked"  help:"The number of tasks that are blocked."`
		Running  int `json:"running"  help:"The number of tasks that are running."`
		Sleeping int `json:"sleeping" help:"The number of tasks that are sleeping."`
		Stopped  int `json:"stopped"  help:"The number of tasks that are stopped."`
		Total    int `json:"total"    help:"The total number of tasks."`
		Zombie   int `json:"zombie"   help:"The number of child tasks that are inactive with an active parent task."`
	} `json:"tasks"`
}

func parseOSMetrics(b []byte) (*osMetrics, error) {
	var m osMetrics
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *osMetrics) originalMetrics(namePrefix, region string) []prometheus.Metric {
	res := make([]prometheus.Metric, 0, 100)
	constLabels := prometheus.Labels{
		"instance": m.InstanceID,
		"region":   region,
	}

	makeMetric := func(name string, help string, labelKeys, labelValues []string, value reflect.Value) {
		var f float64
		switch kind := value.Kind(); kind {
		case reflect.Float64:
			f = value.Float()
		case reflect.Int:
			f = float64(value.Int())
		default:
			panic(fmt.Errorf("can't make a metric value from %v (%s)", value, kind))
		}

		desc := prometheus.NewDesc(namePrefix+"_"+name, help, labelKeys, constLabels)
		metric := prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, f, labelValues...)
		res = append(res, metric)
	}

	// CPUUtilization
	{
		t := reflect.TypeOf(m.CPUUtilization)
		v := reflect.ValueOf(m.CPUUtilization)
		for i := 0; i < t.NumField(); i++ {
			tags := t.Field(i).Tag
			name, help := tags.Get("json"), tags.Get("help")
			makeMetric("cpuUtilization_"+name, help, nil, nil, v.Field(i))
		}
	}

	// DiskIO - move device name to label
	{
		for _, disk := range m.DiskIO {
			labelKeys := []string{"device"}
			labelValues := []string{disk.Device}
			t := reflect.TypeOf(disk)
			v := reflect.ValueOf(disk)
			for i := 0; i < t.NumField(); i++ {
				tags := t.Field(i).Tag
				name, help := tags.Get("json"), tags.Get("help")
				if name == "device" {
					continue
				}
				makeMetric("diskIO_"+name, help, labelKeys, labelValues, v.Field(i))
			}
		}
	}

	// FileSys - move name and mount point to labels
	{
		for _, fs := range m.FileSys {
			labelKeys := []string{"name", "mount_point"}
			labelValues := []string{fs.Name, fs.MountPoint}
			t := reflect.TypeOf(fs)
			v := reflect.ValueOf(fs)
			for i := 0; i < t.NumField(); i++ {
				tags := t.Field(i).Tag
				name, help := tags.Get("json"), tags.Get("help")
				switch name {
				case "name", "mountPoint":
					continue
				}
				makeMetric("fileSys_"+name, help, labelKeys, labelValues, v.Field(i))
			}
		}
	}

	// LoadAverageMinute
	{
		t := reflect.TypeOf(m.LoadAverageMinute)
		v := reflect.ValueOf(m.LoadAverageMinute)
		for i := 0; i < t.NumField(); i++ {
			tags := t.Field(i).Tag
			name, help := tags.Get("json"), tags.Get("help")
			makeMetric("loadAverageMinute_"+name, help, nil, nil, v.Field(i))
		}
	}

	// Memory
	{
		t := reflect.TypeOf(m.Memory)
		v := reflect.ValueOf(m.Memory)
		for i := 0; i < t.NumField(); i++ {
			tags := t.Field(i).Tag
			name, help := tags.Get("json"), tags.Get("help")
			makeMetric("memory_"+name, help, nil, nil, v.Field(i))
		}
	}

	// Network - move interface name to label
	{
		for _, n := range m.Network {
			labelKeys := []string{"interface"}
			labelValues := []string{n.Interface}
			t := reflect.TypeOf(n)
			v := reflect.ValueOf(n)
			for i := 0; i < t.NumField(); i++ {
				tags := t.Field(i).Tag
				name, help := tags.Get("json"), tags.Get("help")
				if name == "interface" {
					continue
				}
				makeMetric("network_"+name, help, labelKeys, labelValues, v.Field(i))
			}
		}
	}

	// ProcessList - move process name, ID, parent ID, thread ID to labels
	{
		for _, p := range m.ProcessList {
			labelKeys := []string{"name", "id", "parentID", "tgid"}
			labelValues := []string{p.Name, strconv.Itoa(p.ID), strconv.Itoa(p.ParentID), strconv.Itoa(p.TGID)}
			t := reflect.TypeOf(p)
			v := reflect.ValueOf(p)
			for i := 0; i < t.NumField(); i++ {
				tags := t.Field(i).Tag
				name, help := tags.Get("json"), tags.Get("help")
				switch name {
				case "name", "id", "parentID", "tgid":
					continue
				}
				makeMetric("processList_"+name, help, labelKeys, labelValues, v.Field(i))
			}
		}
	}

	// Swap
	{
		t := reflect.TypeOf(m.Swap)
		v := reflect.ValueOf(m.Swap)
		for i := 0; i < t.NumField(); i++ {
			tags := t.Field(i).Tag
			name, help := tags.Get("json"), tags.Get("help")
			makeMetric("swap_"+name, help, nil, nil, v.Field(i))
		}
	}

	// Tasks
	{
		t := reflect.TypeOf(m.Tasks)
		v := reflect.ValueOf(m.Tasks)
		for i := 0; i < t.NumField(); i++ {
			tags := t.Field(i).Tag
			name, help := tags.Get("json"), tags.Get("help")
			makeMetric("tasks_"+name, help, nil, nil, v.Field(i))
		}
	}

	return res
}
