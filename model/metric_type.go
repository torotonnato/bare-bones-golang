package model

const (
	TYPE_UNSPECIFIED int32 = 0
	TYPE_COUNT       int32 = 1
	TYPE_RATE        int32 = 2
	TYPE_GAUGE       int32 = 3
)

func IsMetricTypeValid(t int32) bool {
	return t >= TYPE_UNSPECIFIED && t <= TYPE_GAUGE
}

type MetricTypeError struct{}

func (e MetricTypeError) Error() string {
	return "metric: invalid type"
}
