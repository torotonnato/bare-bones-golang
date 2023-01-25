package model

const (
	TYPE_UNSPECIFIED = 0
	TYPE_COUNT       = 1
	TYPE_RATE        = 2
	TYPE_GAUGE       = 3
)

type MetricTypeError struct{}

func (e MetricTypeError) Error() string {
	return "metric: invalid type"
}
