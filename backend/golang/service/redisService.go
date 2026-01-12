package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

var expirationTime time.Duration = time.Hour

type RedisService struct {
	logger  *slog.Logger
	Client  *redis.Client
	Context context.Context
}

func NewRedisService(logger *slog.Logger, addr string, password string, db int) *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	ctx := context.Background()

	return &RedisService{
		logger:  logger.With("redisService"),
		Client:  rdb,
		Context: ctx,
	}
}

func (s *RedisService) GetRandomKey() string {
	key, err := s.Client.RandomKey(s.Context).Result()
	if err != nil {
		s.logger.Error("Failed to get random key: ", err)
		panic(err)
	}
	return key
}

func (s *RedisService) GetBytes(key string) ([]byte, error) {
	val, err := s.Client.Get(s.Context, key).Bytes()
	if err != nil {
		panic(err)
	}
	return val, nil
}

func (s *RedisService) SetBytes(key string, val []byte) {
	s.Client.Set(s.Context, key, val, expirationTime).Err()
}
