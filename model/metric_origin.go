package model

type Origin struct {
	MetricType int32 `json:"metric_type,omitempty"`
	Product    int32 `json:"product,omitempty"`
	Service    int32 `json:"service,omitempty"`
}
