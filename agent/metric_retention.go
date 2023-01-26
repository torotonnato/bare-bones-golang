package agent

import (
	"sort"

	"github.com/torotonnato/gobarebones/api"
	"github.com/torotonnato/gobarebones/config"
	"github.com/torotonnato/gobarebones/model"
)

type MetricItem struct {
	ID model.MetricID
	model.Point
}

type MetricsAccBuffer struct {
	buffer []MetricItem
}

func (a *MetricsAccBuffer) Len() int {
	return len(a.buffer)
}

func (a *MetricsAccBuffer) Swap(i, j int) {
	a.buffer[i], a.buffer[j] = a.buffer[j], a.buffer[i]
}

func (a *MetricsAccBuffer) Less(i, j int) bool {
	return a.buffer[i].ID < a.buffer[j].ID
}

func (a *MetricsAccBuffer) Clear() {
	a.buffer = nil
}

func (a *MetricsAccBuffer) Accumulate(m *MetricItem) {
	a.buffer = append(a.buffer, *m)
}

func (a *MetricsAccBuffer) PastLimit() bool {
	limit := config.AgentMinMetricElementsPerSeries
	return len(a.buffer) >= limit
}

func (a *MetricsAccBuffer) ToSeries() *model.Series {
	if len(a.buffer) == 0 {
		return nil
	}
	sort.Sort(a)
	s := model.NewSeries()
	lastID := model.InvalidMetricID
	currIdx := -1
	for _, mi := range a.buffer {
		if mi.ID != lastID {
			currIdx += 1
			s.Append(regMetrics[mi.ID])
			lastID = mi.ID
		}
		s.Series[currIdx].AppendPoint(mi.Point)
	}
	return s
}

func (a *MetricsAccBuffer) Send() error {
	s := a.ToSeries()
	if s != nil {
		err := api.Series(s)
		a.Clear()
		return err
	}
	return nil
}
