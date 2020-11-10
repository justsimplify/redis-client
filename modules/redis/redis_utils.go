package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/justsimplify/redis-client/modules"
	"log"
)

var redisClient *redis.Client

func getClient(r Client) *redis.Client {
	redisHost := r.Host
	redisPort := r.Port
	redisPassword := r.Password
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     redisHost + ":" + redisPort,
			Password: redisPassword, // no password set
			DB:       0,  // use default DB
		})
	}
	_, err := redisClient.Ping(modules.GenerateContext()).Result()
	if err != nil {
		log.Fatalf("redis client error: %s", err)
	}
	return redisClient
}
