package redis

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/redis/go-redis/v9"
)

type redisConfig struct {
	Password string `env:"REDIS_PASSWORD"`
}

func loadRedisConfig() (redisConfig, error) {
	cfg := redisConfig{}

	if err := env.Parse(&cfg); err != nil {
		return redisConfig{}, err
	}

	emptyConfig := redisConfig{}
	if cfg == emptyConfig {
		return redisConfig{}, errors.New("[loadRedisConfig] redis config is empty")
	}

	log.Printf("redis config successfully loaded")

	return cfg, nil
}

func newRedisConfigURL(p redisConfig) string {
	return fmt.Sprintf("redis://default:%s@redis_service/0?protocol=3",
		p.Password,
	)
}

func LoadRedisClient() (*redis.Client, error) {
	cfg, err := loadRedisConfig()
	if err != nil {
		return nil, err
	}

	opts, err := redis.ParseURL(newRedisConfigURL(cfg))
	if err != nil {
		return nil, fmt.Errorf("[LoadRedisClient] failed to parse config: %w", err)
	}

	redisClient := redis.NewClient(opts)

	if err = testPing(redisClient); err != nil {
		return nil, fmt.Errorf("[LoadRedisClient] failed to ping redis: %w", err)
	}

	return redisClient, nil
}

func testPing(redisClient *redis.Client) error {
	if err := redisClient.Ping(context.Background()); err != nil {
		return err.Err()
	}

	return nil
}
