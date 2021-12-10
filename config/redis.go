package config

import (
	"luwjistik/exception"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(configuration Config) *redis.Client {
	redisDatabase, err := strconv.Atoi(configuration.Get("REDIS_DATABASE"))
	exception.PanicIfNeeded(err)

	db := redis.NewClient(&redis.Options{
		Addr:     configuration.Get("REDIS_URI"),
		Password: configuration.Get("REDIS_PASSWORD"),
		DB:       redisDatabase,
	})
	return db
}
