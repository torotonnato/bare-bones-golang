package agent

import (
	"sync"

	"github.com/torotonnato/gobarebones/config"
	rb "github.com/torotonnato/gobarebones/ringbufferchan"
)

const (
	agentStop = iota
	agentFlush
)

var state struct {
	isRunning bool
	dataChan  *rb.RingBufferChan[any]
	cmdChan   *rb.RingBufferChan[int]
	sync.WaitGroup
	sync.Mutex
}

func worker() {
	metricsBuff := MetricsAccBuffer{}
	for {
		select {
		case item := <-state.dataChan.ReadChan:
			switch data := item.(type) {
			case MetricItem:
				metricsBuff.Accumulate(&data)
				if metricsBuff.PastLimit() {
					metricsBuff.Send()
					tickerReset()
				}
			}
		case cmd := <-state.cmdChan.ReadChan:
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
		return Error{AgentAlreadyRunning}
	}
	capacity := config.AgentChannelCapacity
	state.dataChan = rb.NewRingBufferChan[any](capacity)
	state.cmdChan = rb.NewRingBufferChan[int](capacity)
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
		return Error{AgentNotRunning}
	}
	tickerReset()
	state.cmdChan.WriteChan <- agentFlush
	return nil
}

func Stop() error {
	state.Lock()
	defer state.Unlock()
	if !state.isRunning {
		return Error{AgentNotRunning}
	}
	state.cmdChan.WriteChan <- agentStop
	state.Wait()
	tickerStop()
	state.dataChan.Close()
	state.cmdChan.Close()
	state.isRunning = false
	return nil
}

func FlushAndStop() error {
	if err := Flush(); err != nil {
		return err
	}
	return Stop()
}
