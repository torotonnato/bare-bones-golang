package agent

import (
	"time"

	"github.com/torotonnato/gobarebones/model"
)

func PushMetric(m *model.Metric, value float64) error {
	if !agent.isRunning {
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
