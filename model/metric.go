package model

import (
	"encoding/json"
	"errors"
)

type Metric struct {
	Interval       int64     `json:"interval,omitempty"`
	Metadata       *Metadata `json:"metadata,omitempty"`
	Metric         string    `json:"metric"`
	Points         []Point   `json:"points"`
	Resources      *Resource `json:"resources,omitempty"`
	SourceTypeName string    `json:"source_type_name,omitempty"`
	Tags           Tags      `json:"tags,omitempty"`
	Type           int32     `json:"type"`
	Unit           string    `json:"unit,omitempty"`
}

func NewMetric(name string, t int32) (*Metric, error) {
	if t < TYPE_UNSPECIFIED || t > TYPE_GAUGE {
		return nil, errors.New("invalid type")
	}
	m := &Metric{}
	m.Metric = name
	m.Type = t
	return m, nil
}

func (m *Metric) SetInterval(interval int64) *Metric {
	if m.Type != TYPE_RATE && m.Type != TYPE_COUNT {
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

func (m *Metric) SetMetadata(metadata *Metadata) *Metric {
	m.Metadata = metadata
	return m
}

func (m *Metric) SetResources(resources *Resource) *Metric {
	m.Resources = resources
	return m
}

func (m *Metric) SetSourceTypeName(sourceTypeName string) *Metric {
	m.SourceTypeName = sourceTypeName
	return m
}

func (m *Metric) SetTags(tags Tags) *Metric {
	m.Tags = tags
	return m
}

func (m *Metric) SetUnit(unit string) *Metric {
	m.Unit = unit
	return m
}

func (m *Metric) toJSON() ([]byte, error) {
	return json.Marshal(m)
}

/*
func (m *Metric) Sample(val float64) {
	p := Point{
		time.Now().Unix(),
		val,
	}
	sample := agentSample{
		m:           m,
		metricPoint: p,
	}
	agentChannel <- sample
}
*/
