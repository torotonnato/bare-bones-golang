package agent

const (
	MetricAlreadyExists = iota + 1
)

type Error struct {
	Code int
}

func (e Error) Error() string {
	switch e.Code {
	case MetricAlreadyExists:
		return "agent: metric already registered"
	}
	return "agent: unknown error"
}
