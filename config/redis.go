package config

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(configuration Config) *redis.Client {

	db := redis.NewClient(&redis.Options{
		Addr:     configuration.Get("REDIS_URI"),
		Password: configuration.Get("REDIS_PASSWORD"),
		Username: configuration.Get("REDIS_USERNAME"),
	})
	return db
}
