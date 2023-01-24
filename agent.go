package gobarebones

import (
	"fmt"
	"log"
)

type agentSample struct {
	m *metric
	metricPoint
}

const agentChannelSize = 128

var agentChannel = make(chan agentSample, agentChannelSize)

func agent() {
	log.Println("Agent started...")
	samples := make(map[string][]metricPoint)
	metrics := make(map[string]*metric)
	for s := range agentChannel {
		name := s.m.Metric
		if _, ok := samples[name]; !ok {
			samples[name] = make([]metricPoint, 16)
		}
		samples[name] = append(samples[name], s.metricPoint)
		metrics[name] = s.m

		fmt.Println(name, s.metricPoint)
	}
}

func init() {
	go agent()
}
