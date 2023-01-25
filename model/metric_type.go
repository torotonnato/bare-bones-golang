package model

type MetricType int32

const (
	TYPE_UNSPECIFIED MetricType = 0
	TYPE_COUNT       MetricType = 1
	TYPE_RATE        MetricType = 2
	TYPE_GAUGE       MetricType = 3
)

func (t MetricType) IsValid() bool {
	return t >= TYPE_UNSPECIFIED && t <= TYPE_GAUGE
}

func (e MetricType) NeedsInterval() bool {
	return e == TYPE_RATE || e == TYPE_COUNT
}

type MetricTypeError struct{}

func (e MetricTypeError) Error() string {
	return "metric: invalid type"
}
