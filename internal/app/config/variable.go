package config

type ObservationVariables struct {
	JaegerEndpoint string `env:"JAEGER_ENDPOINT"`
}

type Variables struct {
	Port        uint16 `env:"PORT"`
	Environment string `env:"ENVIRONMENT"`
	Observation ObservationVariables
}
