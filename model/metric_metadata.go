package model

type Metadata struct {
	Origin Origin `json:"origin,omitempty"`
}

func NewMetricMetadata(mtype int32, product int32, service int32) *Metadata {
	origin := Origin{
		MetricType: mtype,
		Product:    product,
		Service:    service,
	}
	return &Metadata{
		Origin: origin,
	}
}
