package agent

import (
	"sync"

	"github.com/torotonnato/gobarebones/config"
)

const (
	agentStop = iota
	agentFlush
)

type state struct {
	isRunning bool
	channel   chan interface{}
	sync.WaitGroup
}

var agent state

func worker() {
	metricsBuff := MetricsAccBuffer{}
	for item := range agent.channel {
		switch i := item.(type) {
		case MetricItem:
			metricsBuff.Accumulate(&i)
			if metricsBuff.PastLimit() {
				metricsBuff.Send()
			}
		case int:
			if i == agentStop {
				agent.Done()
				return
			} else if i == agentFlush {
				metricsBuff.Send()
			}
		}
	}
}

func Start() error {
	if agent.isRunning {
		return Error{Code: AgentAlreadyRunning}
	}
	agent.isRunning = true
	agent.channel = make(chan interface{}, config.AgentChannelCapacity)
	agent.Add(1)
	go worker()
	return nil
}

func Flush() error {
	if !agent.isRunning {
		return Error{Code: AgentNotRunning}
	}
	agent.channel <- agentFlush
	return nil
}

func Stop() error {
	if !agent.isRunning {
		return Error{Code: AgentNotRunning}
	}
	agent.channel <- agentStop
	agent.Wait()
	agent.isRunning = false
	return nil
}

func FlushAndStop() error {
	if err := Flush(); err != nil {
		return err
	}
	return Stop()
}
