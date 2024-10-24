package sessions

import "github.com/redis/go-redis/v9"

type SessionRepo struct {
	redis *redis.Client
}

func NewSessionRepo(redisClient *redis.Client) *SessionRepo {
	return &SessionRepo{
		redis: redisClient,
	}
}
