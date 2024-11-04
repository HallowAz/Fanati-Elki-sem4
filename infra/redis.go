package infra

import (
	"fe-sem4/config"
	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: config.RedisHost + config.RedisPort,
	})
}

func ConnectToRedis(cli *redis.Client) error {
	return cli.Ping().Err()
}
