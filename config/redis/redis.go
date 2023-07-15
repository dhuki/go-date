package redis

import (
	"github.com/dhuki/go-date/config"
	"github.com/go-redis/redis"
)

func InitRedis(conf *config.RedisConfig) (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Password: conf.Password,
		DB:       conf.DB,
	})

	if _, err = RedisClient.Ping().Result(); err != nil {
		return
	}
	return
}
