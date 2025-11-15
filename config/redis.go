package config

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/google/wire"
)

var RedisProviderSet = wire.NewSet(NewRedisClient)

func NewRedisClient(AppConfig *Config) *redis.Client {
	if AppConfig.Redis.Addr != "" {
		RedisClient := redis.NewClient(&redis.Options{
			Addr:     AppConfig.Redis.Addr,
			Password: AppConfig.Redis.Password,
			DB:       AppConfig.Redis.DB,
		})
		//这个是ping 是否连接成功
		_, err := RedisClient.Ping().Result()
		if err != nil {
			log.Fatalf("failed to connect redis: %v", err)
		}
		return RedisClient
	}
	return nil
}
