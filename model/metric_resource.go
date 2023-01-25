package model

type Resource struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

func NewResource(name string, mtype string) *Resource {
	return &Resource{
		Name: name,
		Type: mtype,
	}
}
