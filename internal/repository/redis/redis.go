package redis

import (
	"context"
	"fmt"
	"github.com/blazee5/finance-tracker/internal/config"
	"github.com/redis/go-redis/v9"
)

func Run(cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: "",
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return rdb
}
