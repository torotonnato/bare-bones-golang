package model

const (
	InvalidMetricType = iota
)

type Error struct {
	Code int
}

func (e Error) Error() string {
	switch e.Code {
	case InvalidMetricType:
		return "invalid metric type"
	}
	return "unknown error"
}
