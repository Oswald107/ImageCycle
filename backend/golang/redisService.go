package main

import (
	"context"
	"log"

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
	key, err := rdb.RandomKey(ctx).Result()
	if err != nil {
		log.Fatal("Failed to get random key")
		panic(err)
	}
	return key
}

func (s *RedisService) GetBytes() ([]byte, error) {
	val, err := rdb.Get(ctx, key).Bytes()
	if err != nil {
		panic(err)
	}
	return val
}

func (s *RedisService) SetBytes() {
	return s.Client.Set(s.Context, key, value, 0).Err()
}
