package model

import "sync/atomic"

var globalMetricId int32 = 0

func GetUniqueId() int32 {
	return atomic.AddInt32(&globalMetricId, 1)
}
