package enhanced2

import (
	"strings"
)

var (
	dataMySQL57 = []byte(strings.TrimSpace(`
	{
		"engine": "MYSQL",
		"instanceID": "rds-mysql57",
		"instanceResourceID": "db-EXAMPLE",
		"timestamp": "2018-01-08T14:35:21Z",
		"version": 1.00,
		"uptime": "39 days, 5:08:13",
		"numVCPUs": 1,
		"cpuUtilization": {
			"guest": 0.00,
			"irq": 1.16,
			"system": 4.71,
			"wait": 81.43,
			"idle": 0.00,
			"user": 1.18,
			"total": 99.99,
			"steal": 1.01,
			"nice": 10.50
		},
		"loadAverageMinute": {
			"fifteen": 7.64,
			"five": 14.08,
			"one": 32.83
		},
		"memory": {
			"writeback": 0,
			"hugePagesFree": 0,
			"hugePagesRsvd": 0,
			"hugePagesSurp": 0,
			"cached": 162452,
			"hugePagesSize": 2048,
			"free": 127048,
			"hugePagesTotal": 0,
			"inactive": 511232,
			"pageTables": 7916,
			"dirty": 360,
			"mapped": 14184,
			"active": 1299668,
			"total": 2051520,
			"slab": 45440,
			"buffers": 149412
		},
		"tasks": {
			"sleeping": 258,
			"zombie": 0,
			"running": 1,
			"stopped": 0,
			"total": 313,
			"blocked": 54
		},
		"swap": {
			"cached": 908,
			"total": 4095996,
			"out": 0.00,
			"free": 3803212,
			"in": 0.00
		},
		"network": [
			{
				"interface": "eth0",
				"rx": 93611.95,
				"tx": 801373.55
			}
		],
		"diskIO": [
			{
				"writeKbPS": 6157.27,
				"readIOsPS": 354.18,
				"await": 74.67,
				"readKbPS": 5664.93,
				"rrqmPS": 0.00,
				"util": 99.94,
				"avgQueueLen": 2783.31,
				"tps": 623.33,
				"readKb": 339896,
				"device": "rdsdev",
				"writeKb": 369436,
				"avgReqSz": 18.97,
				"wrqmPS": 0.00,
				"writeIOsPS": 269.15
			}
		],
		"fileSys": [
			{
				"used": 9659328,
				"name": "rdsfilesys",
				"usedFiles": 731,
				"usedFilePercent": 0.02,
				"maxFiles": 3932160,
				"mountPoint": "/rdsdbdata",
				"total": 61774768,
				"usedPercent": 15.64
			}
		],
		"processList": [
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1594,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1595,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1596,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1597,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1599,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1600,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1601,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1602,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.10,
				"id": 1603,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1604,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1605,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1606,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1607,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1608,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1609,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1610,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1611,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1612,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1613,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1614,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1615,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1616,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1617,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1618,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1619,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1620,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1621,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1622,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1623,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1624,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1625,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.15,
				"id": 1626,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.10,
				"id": 1627,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1628,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1629,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1630,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1631,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1632,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1633,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1634,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1635,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1636,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1637,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1638,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.12,
				"id": 1639,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.15,
				"id": 1640,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1641,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1642,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.13,
				"id": 1643,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.00,
				"id": 4516,
				"rss": 1282364
			},
			{
				"vss": 2552316,
				"name": "mysqld",
				"tgid": 4516,
				"vmlimit": "unlimited",
				"parentID": 1,
				"memoryUsedPc": 62.51,
				"cpuUsedPc": 0.33,
				"id": 4551,
				"rss": 1282364
			},
			{
				"vss": 723152,
				"name": "OS processes",
				"tgid": 0,
				"vmlimit": "",
				"parentID": 0,
				"memoryUsedPc": 0.53,
				"cpuUsedPc": 0.02,
				"id": 0,
				"rss": 10968
			},
			{
				"vss": 1307432,
				"name": "RDS processes",
				"tgid": 0,
				"vmlimit": "",
				"parentID": 0,
				"memoryUsedPc": 10.58,
				"cpuUsedPc": 0.14,
				"id": 0,
				"rss": 217304
			}
		]
	}`))
)
