package gobarebones

import (
	"fmt"
	"log"

	"github.com/torotonnato/gobarebones/model"
)

type agentSample struct {
	m *model.Metric
	model.Point
}

const agentChannelSize = 128

var agentChannel = make(chan agentSample, agentChannelSize)

func agent() {
	log.Println("Agent started...")
	samples := make(map[string][]model.Point)
	metrics := make(map[string]*model.Metric)
	for s := range agentChannel {
		name := s.m.Metric
		if _, ok := samples[name]; !ok {
			samples[name] = make([]model.Point, 16)
		}
		samples[name] = append(samples[name], s.Point)
		metrics[name] = s.m

		fmt.Println(name, s.Point)
	}
}

func init() {
	go agent()
}
