package redis

type Config struct {
	Address       string `env:"REDIS_ADDRESS" default:"localhost:6379"`
	EnableTracing bool   `env:"ENABLE_TRACING" default:"false"`
}
