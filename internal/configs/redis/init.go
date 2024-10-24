package redis

import (
	"context"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/redis/go-redis/v9"
	"log"
)

type redisConfig struct {
	User     string `env:"REDIS_USER"`
	Password string `env:"REDIS_PASSWORD"`
}

func mustLoadRedisConfig() redisConfig {
	cfg := redisConfig{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	test := redisConfig{}
	if test == cfg {
		log.Fatal("[mustLoadRedisConfig] redis config is empty")
	}

	log.Printf("redis config successfully loaded")

	return cfg
}

func newRedisConfigURL(p redisConfig) string {
	return fmt.Sprintf("redis://default:%s@redis_service/0?protocol=3",
		p.Password,
	)
}

func MustLoadRedisClient() *redis.Client {
	cfg := mustLoadRedisConfig()

	opts, err := redis.ParseURL(newRedisConfigURL(cfg))
	if err != nil {
		log.Fatalf("[MustLoadRedisClient] failed to parse config: %s", err.Error())
	}

	redisClient := redis.NewClient(opts)

	if err = testPing(redisClient); err != nil {
		log.Fatalf("[TestPing] failed to ping redis: %s", err.Error())
	}

	return redisClient
}

func testPing(redisClient *redis.Client) error {
	if err := redisClient.Ping(context.Background()); err != nil {
		return err.Err()
	}

	return nil
}
