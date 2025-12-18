package tracer

type Configuration struct {
	CollectorEndpoint           string  `env:"COLLECTOR_ENDPOINT"`
	ServiceName                 string  `env:"SERVICE_NAME"`
	ExportBatchSize             int     `env:"EXPORT_BATCH_SIZE"`
	BatchTimeOut                int     `env:"BATCH_TIMEOUT"`
	MaxQueueSize                int     `env:"MAX_QUEUE_SIZE"`
	SampleRatio                 float64 `env:"SAMPLE_RATIO"`
	AttributeCountLimit         int     `env:"ATTRIBUTE_COUNT_LIMIT"`
	EventCountLimit             int     `env:"EVENT_COUNT_LIMIT"`
	LinkCountLimit              int     `env:"LINK_COUNT_LIMIT"`
	AttributePerEventCountLimit int     `env:"ATTRIBUTE_PER_EVENT_COUNT_LIMIT"`
	AttributePerLinkCountLimit  int     `env:"ATTRIBUTE_PER_LINK_COUNT_LIMIT"`
}

func DefaultConfig() *Configuration {
	return &Configuration{
		CollectorEndpoint:           "localhost:4318",
		ServiceName:                 "Hlidskjalf",
		ExportBatchSize:             512,
		BatchTimeOut:                5,
		MaxQueueSize:                2048,
		SampleRatio:                 1.0,
		AttributeCountLimit:         128,
		EventCountLimit:             128,
		LinkCountLimit:              128,
		AttributePerEventCountLimit: 128,
		AttributePerLinkCountLimit:  128,
	}
}
