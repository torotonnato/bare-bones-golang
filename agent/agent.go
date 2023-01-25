package agent

import (
	"github.com/torotonnato/gobarebones/model"
)

var regMetrics = make(map[int32]model.Metric)

func RegisterMetric(m *model.Metric) error {
	if m != nil {
		if _, exists := regMetrics[m.Id]; exists {
			return AlreadyExists{}
		}
		regMetrics[m.Id] = *m.DeepCopy()
	}
	return nil
}
