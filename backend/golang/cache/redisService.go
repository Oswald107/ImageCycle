package cache

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

var expirationTime time.Duration = time.Hour

type RedisService struct {
	logger  *slog.Logger
	client  *redis.Client
	context context.Context
	curKey  string
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
		client:  rdb,
		context: ctx,
		curKey:  "set1",
	}
}

func (rs *RedisService) GetImageName() (string, error) {
	return rs.client.SRandMember(rs.context, rs.curKey).Result()
}

func (rs *RedisService) StoreImageName(imageName string) error {
	return rs.client.SAdd(rs.context, rs.curKey, imageName).Err()
}

func (rs *RedisService) CreateNewSet(imageNames []string) error {
	key := "set1"
	if rs.curKey == "set1" {
		key = "set2"
	}

	err := rs.client.SAdd(rs.context, rs.curKey, imageNames).Err()
	if err != nil {
		return err
	}

	err = rs.client.Expire(rs.context, key, expirationTime).Err()
	if err != nil {
		return err
	}

	rs.curKey = key

	return nil
}
