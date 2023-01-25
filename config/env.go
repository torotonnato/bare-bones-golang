package config

type Config struct {
	Region string
	APIKey string
}

const (
	AgentMinMetricElementsPerSeries  = 100
	AgentAvgDistinctMetricsPerSeries = 10
	AgentAvgPointsPerMetric          = 10
)

const (
	DD_COM     = "https://api.datadoghq.com"
	US3_DD_COM = "https://api.us3.datadoghq.com"
	US5_DD_COM = "https://us5.datadoghq.com"
	DD_EU      = "https://datadoghq.eu"
	DD_GOV     = "https://ddog-gov.com"
)

var Env Config

func Setup(config *Config) {
	Env = *config
}
