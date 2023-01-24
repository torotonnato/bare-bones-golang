package gobarebones

type Config struct {
	Region string
	APIKey string
}

const (
	DD_COM     = "https://api.datadoghq.com"
	US3_DD_COM = "https://api.us3.datadoghq.com"
	US5_DD_COM = "https://us5.datadoghq.com"
	DD_EU      = "https://datadoghq.eu"
	DD_GOV     = "https://ddog-gov.com"
)

var config Config

func Setup(c *Config) error {
	config = *c
	valid, err := apiValidate()
	if valid {
		return nil
	}
	return err
}
