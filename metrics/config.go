package metrics

type Configuration struct {
	CollectorEndpoint string `json:"collector_endpoint" mapstructure:"collector_endpoint" yaml:"collector_endpoint" env:"COLLECTOR_ENDPOINT"`
	ServiceName       string `json:"service_name" mapstructure:"service_name" yaml:"service_name" env:"SERVICE_NAME"`
	ExportInterval    int    `json:"export_interval" mapstructure:"export_interval" yaml:"export_interval" env:"EXPORT_INTERVAL"`
}

func DefaultConfig() *Configuration {
	return &Configuration{
		CollectorEndpoint: "localhost:4318",
		ServiceName:       "shared",
		ExportInterval:    30,
	}
}
