package sessions

import "github.com/redis/go-redis/v9"

type SessionStore struct {
	redis *redis.Client
}

func NewSessionRepo(redisClient *redis.Client) *SessionStore {
	return &SessionStore{
		redis: redisClient,
	}
}
