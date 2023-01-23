package gobarebones

import (
	"fmt"
)

type Config struct {
	APIKey string
	APPKey string
	Host   string
}

var config

func (c *Config) Setup() {
	config = c
}
