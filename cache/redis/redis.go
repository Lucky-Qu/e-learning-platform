package redis

import (
	"e-learning-platform/config"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Config.Redis.Host, config.Config.Redis.Port),
		Password: config.Config.Redis.Password,
	})
}

func AddStringToRedis(key string, value string, expires time.Duration) error {
	err := rdb.Set(key, value, expires).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetStringFromRedis(key string) (string, error) {
	val, err := rdb.Get(key).Result()
	if err != nil {
		return "", err
	}
	if val == "" {
		return "", errors.New("查询结果为空")
	}
	return val, nil
}
