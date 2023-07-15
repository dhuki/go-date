package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

//go:generate mockgen -destination=mocks/mock_redis.go -package=mocks github.com/dhuki/go-date/pkg/redis Redis

type Redis interface {
	Get(key string) (value string)
	SetLockingKey(key string, value interface{}, ttl time.Duration) (err error)
	SetIncr(key string) (count int64)
	Delete(key string) (err error)
	Set(key string, value interface{}, ttl time.Duration) (err error)
}

type RedisImpl struct {
	redisClient *redis.Client
}

func NewRedisLibs(redisClient *redis.Client) Redis {
	return RedisImpl{
		redisClient: redisClient,
	}
}

func (r RedisImpl) Get(key string) string {
	return r.redisClient.Get(key).Val()
}

func (r RedisImpl) SetLockingKey(key string, value interface{}, ttl time.Duration) (err error) {
	success, err := r.redisClient.SetNX(key, value, ttl).Result()
	if err != nil {
		return
	}
	if !success {
		err = errors.New("error multiple keys in redis")
	}
	return
}

func (r RedisImpl) SetIncr(key string) int64 {
	return r.redisClient.Incr(key).Val()
}

func (r RedisImpl) Delete(key string) (err error) {
	val := r.redisClient.Get(key).Val()
	if len(val) > 0 {
		return r.redisClient.Del(key).Err()
	}
	return
}

func (r RedisImpl) Set(key string, value interface{}, ttl time.Duration) error {
	return r.redisClient.Set(key, value, ttl).Err()
}
