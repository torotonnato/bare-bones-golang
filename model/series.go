package model

type Series struct {
	Series []Metric `json:"series"`
}

func NewSeries() *Series {
	return &Series{[]Metric{}}
}

func (s *Series) Append(m Metric) {
	s.Series = append(s.Series, m)
}
