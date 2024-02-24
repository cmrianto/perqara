package redis

import (
	"fmt"
	"strconv"

	"perqara/config"

	"github.com/go-redis/redis/v8"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis.v8"
)

func NewRedisClient(cfg *config.Config) (redis.UniversalClient, error) {
	db, err := strconv.Atoi(cfg.Redis.Database)
	if err != nil {
		return nil, err
	}

	rdb := redistrace.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       db,
	})

	return rdb, nil
}
