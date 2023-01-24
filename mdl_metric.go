package gobarebones

import (
	"encoding/json"
	"errors"
	"time"
)

type metricOrigin struct {
	MetricType int32 `json:"metric_type,omitempty"`
	Product    int32 `json:"product,omitempty"`
	Service    int32 `json:"service,omitempty"`
}

type metricMetadata struct {
	Origin metricOrigin `json:"origin,omitempty"`
}

func NewMetricMetadata(mtype int32, product int32, service int32) *metricMetadata {
	origin := metricOrigin{
		MetricType: mtype,
		Product:    product,
		Service:    service,
	}
	return &metricMetadata{
		Origin: origin,
	}
}

type metricPoint struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type metricResource struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

func NewMetricResource(name string, mtype string) *metricResource {
	return &metricResource{
		Name: name,
		Type: mtype,
	}
}

// Metric.Type enum
const (
	METRIC_UNSPECIFIED = 0
	METRIC_COUNT       = 1
	METRIC_RATE        = 2
	METRIC_GAUGE       = 3
)

type metricTags []string

func NewMetricTags() metricTags {
	return metricTags{}
}

func (t metricTags) Add(tag string) metricTags {
	return append(t, tag)
}

type metric struct {
	Interval       int64           `json:"interval,omitempty"`
	Metadata       *metricMetadata `json:"metadata,omitempty"`
	Metric         string          `json:"metric"`
	Points         []metricPoint   `json:"points"`
	Resources      *metricResource `json:"resources,omitempty"`
	SourceTypeName string          `json:"source_type_name,omitempty"`
	Tags           metricTags      `json:"tags,omitempty"`
	Type           int32           `json:"type"`
	Unit           string          `json:"unit,omitempty"`
}

func NewMetric(name string, t int32) (*metric, error) {
	if t < METRIC_UNSPECIFIED || t > METRIC_GAUGE {
		return nil, errors.New("invalid type")
	}
	m := &metric{}
	m.Metric = name
	m.Type = t
	return m, nil
}

const MaxInterval int64 = 3600

func (m *metric) SetInterval(interval int64) *metric {
	if m.Type != METRIC_RATE && m.Type != METRIC_COUNT {
		m.Interval = 0
		return m
	}
	if interval < 1 {
		m.Interval = 1
	} else if interval > MaxInterval {
		m.Interval = MaxInterval
	} else {
		m.Interval = interval
	}
	return m
}

func (m *metric) SetMetadata(metadata *metricMetadata) *metric {
	m.Metadata = metadata
	return m
}

func (m *metric) SetResources(resources *metricResource) *metric {
	m.Resources = resources
	return m
}

func (m *metric) SetSourceTypeName(sourceTypeName string) *metric {
	m.SourceTypeName = sourceTypeName
	return m
}

func (m *metric) SetTags(tags metricTags) *metric {
	m.Tags = tags
	return m
}

func (m *metric) SetUnit(unit string) *metric {
	m.Unit = unit
	return m
}

func (m *metric) toJSON() ([]byte, error) {
	return json.Marshal(m)
}

func (m *metric) Sample(val float64) {
	p := metricPoint{
		time.Now().Unix(),
		val,
	}
	sample := agentSample{
		m: m,
		metricPoint: p,
	}
	agentChannel <- sample
}
