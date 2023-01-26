package model

type Series struct {
	Series []Metric `json:"series"`
}

func NewSeries() *Series {
	s := &Series{}
	s.Series = []Metric{}
	return s
}

func (s *Series) Append(m Metric) {
	s.Series = append(s.Series, m)
}
