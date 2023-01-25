package model

type Metric struct {
	ID             MetricID   `json:"-"`
	Interval       int64      `json:"interval,omitempty"`
	Metadata       *Metadata  `json:"metadata,omitempty"`
	Metric         string     `json:"metric"`
	Points         []Point    `json:"points"`
	Resources      *Resource  `json:"resources,omitempty"`
	SourceTypeName string     `json:"source_type_name,omitempty"`
	Tags           Tags       `json:"tags,omitempty"`
	Type           MetricType `json:"type"`
	Unit           string     `json:"unit,omitempty"`
}

func NewMetric(name string, mType MetricType) (*Metric, error) {
	if !mType.IsValid() {
		return nil, MetricTypeError{}
	}
	m := &Metric{}
	m.ID = GetUniqueMetricID()
	m.Metric = name
	m.Type = mType
	return m, nil
}

func (m *Metric) DeepCopy() *Metric {
	deepCopy := Metric{
		ID:             m.ID,
		Interval:       m.Interval,
		Metadata:       nil,
		Metric:         m.Metric,
		Points:         []Point{},
		Resources:      nil,
		SourceTypeName: m.SourceTypeName,
		Tags:           m.Tags,
		Type:           m.Type,
		Unit:           m.Unit,
	}
	if m.Metadata != nil {
		deepCopy.Metadata = &Metadata{
			Origin: m.Metadata.Origin,
		}
	}
	if m.Resources != nil {
		deepCopy.Resources = &Resource{
			Name: m.Resources.Name,
			Type: m.Resources.Type,
		}
	}
	return &deepCopy
}

func (m *Metric) Clone() *Metric {
	clone := m.DeepCopy()
	m.ID = GetUniqueMetricID()
	return clone
}

func (m *Metric) SetInterval(interval int64) *Metric {
	if m.Type.NeedsInterval() {
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
