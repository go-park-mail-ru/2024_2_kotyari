package sessions

import (
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type SessionStore struct {
	redis *redis.Client
	log   *slog.Logger
}

func NewSessionRepo(redisClient *redis.Client, log *slog.Logger) *SessionStore {
	return &SessionStore{
		redis: redisClient,
		log:   log,
	}
}
