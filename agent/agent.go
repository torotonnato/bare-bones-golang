package agent

import (
	"sync"

	"github.com/torotonnato/gobarebones/config"
)

const (
	agentStop = iota
	agentFlush
)

var state struct {
	isRunning bool
	dataChan  chan interface{}
	cmdChan   chan int
	sync.WaitGroup
	sync.Mutex
}

func blackhole(x interface{}) {}

func drainChannels() {
	for item := range state.dataChan {
		blackhole(item)
	}
	for cmd := range state.cmdChan {
		blackhole(cmd)
	}
}

func worker() {
	metricsBuff := MetricsAccBuffer{}
	for {
		select {
		case item := <-state.dataChan:
			switch data := item.(type) {
			case MetricItem:
				metricsBuff.Accumulate(&data)
				if metricsBuff.PastLimit() {
					metricsBuff.Send()
					tickerReset()
				}
			}
		case cmd := <-state.cmdChan:
			if cmd == agentStop {
				state.Done()
				return
			} else if cmd == agentFlush {
				metricsBuff.Send()
			}
		case <-beat.ticker.C:
			metricsBuff.Send()
		}
	}
}

func Start() error {
	state.Lock()
	defer state.Unlock()
	if state.isRunning {
		return Error{Code: AgentAlreadyRunning}
	}
	state.dataChan = make(chan interface{}, config.AgentChannelCapacity)
	state.cmdChan = make(chan int, config.AgentChannelCapacity)
	state.Add(1)
	tickerStart()
	go worker()
	state.isRunning = true
	return nil
}

func Flush() error {
	state.Lock()
	defer state.Unlock()
	if !state.isRunning {
		return Error{Code: AgentNotRunning}
	}
	tickerReset()
	state.cmdChan <- agentFlush
	return nil
}

func Stop() error {
	state.Lock()
	defer state.Unlock()
	if !state.isRunning {
		return Error{Code: AgentNotRunning}
	}
	state.cmdChan <- agentStop
	state.Wait()
	tickerStop()
	drainChannels()
	state.isRunning = false
	return nil
}

func FlushAndStop() error {
	if err := Flush(); err != nil {
		return err
	}
	return Stop()
}
