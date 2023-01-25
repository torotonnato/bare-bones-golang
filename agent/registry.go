package agent

import (
	"github.com/torotonnato/gobarebones/model"
)

var regMetrics = make(map[model.MetricID]model.Metric)

func RegisterMetric(m *model.Metric) error {
	if m != nil {
		if _, exists := regMetrics[m.ID]; exists {
			return Error{Code: MetricAlreadyExists}
		}
		regMetrics[m.ID] = *m.DeepCopy()
	}
	return nil
}
