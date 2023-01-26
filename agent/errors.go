package agent

const (
	MetricAlreadyExists = iota + 1
	AgentNotRunning
	AgentAlreadyRunning
)

type Error struct {
	Code int
}

func (e Error) Error() string {
	switch e.Code {
	case MetricAlreadyExists:
		return "metric already registered"
	}
	return "unknown error"
}
