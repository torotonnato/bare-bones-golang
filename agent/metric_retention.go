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
	container []MetricItem
}

func (a *MetricsAccBuffer) Len() int {
	return len(a.container)
}

func (a *MetricsAccBuffer) Swap(i, j int) {
	a.container[i], a.container[j] = a.container[j], a.container[i]
}

func (a *MetricsAccBuffer) Less(i, j int) bool {
	return a.container[i].ID < a.container[j].ID
}

func (a *MetricsAccBuffer) Clear() {
	a.container = nil
}

func (a *MetricsAccBuffer) Accumulate(m *MetricItem) {
	a.container = append(a.container, *m)
}

func (a *MetricsAccBuffer) PastLimit() bool {
	return len(a.container) >= config.AgentMinMetricElementsPerSeries
}

func (a *MetricsAccBuffer) ToSeries() *model.Series {
	if len(a.container) == 0 {
		return nil
	}
	sort.Sort(a)
	s := model.Series{}
	s.Series = make([]model.Metric, 0, config.AgentAvgDistinctMetricsPerSeries)
	lastID := model.InvalidMetricID
	currIdx := -1
	for _, p := range a.container {
		if p.ID != lastID {
			currIdx += 1
			s.Series = append(s.Series, regMetrics[p.ID])
			s.Series[currIdx].Points = make([]model.Point, 0, config.AgentAvgPointsPerMetric)
			lastID = p.ID
		}
		s.Series[currIdx].Points = append(s.Series[currIdx].Points, p.Point)
	}
	return &s
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
