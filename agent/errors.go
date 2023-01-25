package agent

type AlreadyExists struct{}

func (e AlreadyExists) Error() string {
	return "metric already registered"
}
