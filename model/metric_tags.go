package model

type Tags []string

func NewTags() Tags {
	return Tags{}
}

func (t Tags) Add(tag string) Tags {
	return append(t, tag)
}