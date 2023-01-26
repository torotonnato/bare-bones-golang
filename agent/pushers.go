package agent

import (
	"time"

	"github.com/torotonnato/gobarebones/model"
)

func PushMetric(m *model.Metric, value float64) error {
	if !state.isRunning {
		return Error{Code: AgentNotRunning}
	}
	item := MetricItem{
		ID: m.ID,
		Point: model.Point{
			Value:     value,
			Timestamp: time.Now().Unix(),
		},
	}
	state.dataChan <- item
	return nil
}
