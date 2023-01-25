package model

import "sync/atomic"

type MetricID = int32

const InvalidMetricID = MetricID(-1)

var globalMetricID MetricID = 0

func GetUniqueMetricID() MetricID {
	return atomic.AddInt32(&globalMetricID, 1)
}
