package redis

import (
	"github.com/YumikoKawaii/shared/logger"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

func Initialize(cfg Config) *redis.Client {

	client := redis.NewClient(
		&redis.Options{
			Addr: cfg.Address,
		},
	)
	if cfg.EnableTracing {
		if err := redisotel.InstrumentTracing(client); err != nil {
			logger.Errorf("error enable tracing redis: %s", err.Error())
		}
	}
	return client
}
