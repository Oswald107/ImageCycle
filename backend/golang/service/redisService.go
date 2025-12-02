package service

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	Client  *redis.Client
	Context context.Context
}

func NewRedisService(addr string, password string, db int) *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	ctx := context.Background()

	return &RedisService{
		Client:  rdb,
		Context: ctx,
	}
}

func (s *RedisService) GetRandomKey() string {
	key, err := s.Client.RandomKey(s.Context).Result()
	if err != nil {
		log.Fatal("Failed to get random key")
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
	s.Client.Set(s.Context, key, val, time.Second*10).Err()
}
