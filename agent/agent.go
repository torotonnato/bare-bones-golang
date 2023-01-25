package agent

import (
	"log"
	"time"

	"sync"

	"github.com/torotonnato/gobarebones/model"
)

const (
	agentStop = iota
	agentFlush
)

type state struct {
	channel chan interface{}
	sync.WaitGroup
	running bool
}

var agent state

func PushMetric(m *model.Metric, value float64) error {
	if !agent.running {
		return Error{Code: AgentNotRunning}
	}
	item := MetricItem{
		From: m.ID,
		Point: model.Point{
			Value:     value,
			Timestamp: time.Now().Unix(),
		},
	}
	agent.channel <- item
	return nil
}

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
	log.Println("Agent started")
	if agent.running {
		return Error{Code: AgentAlreadyRunning}
	}
	agent.running = true
	agent.channel = make(chan interface{}, 256)
	agent.Add(1)
	go worker()
	return nil
}

func Flush() error {
	log.Println("Agent flush")
	if !agent.running {
		return Error{Code: AgentNotRunning}
	}
	agent.channel <- agentFlush
	return nil
}

func Stop() error {
	log.Println("Agent stopped")
	if !agent.running {
		return Error{Code: AgentNotRunning}
	}
	agent.channel <- agentStop
	agent.Wait()
	agent.running = false
	return nil
}
