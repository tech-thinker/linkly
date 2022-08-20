package database

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

var (
	// IsRedisConnected returns the connection status
	IsRedisConnected bool
	Redis            *redis.Client
)

// GetRedis returns the redis client
func GetRedis() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	if Redis == nil {
		if redisHost == "" {
			fmt.Println("Environment variable REDIS_HOST is null.")
			return nil
		}
		if redisPort == "" {
			fmt.Println("Environment variable REDIS_PORT is null.")
			return nil
		}
		if redisPassword == "" {
			fmt.Println("Environment variable REDIS_PASSWORD is null.")
			return nil
		}
	}

	DB := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
	})
	_, err := DB.Ping().Result()
	if err == nil {
		IsRedisConnected = true
	} else {
		log.Println("failed to connect redis")
	}

	return Redis
}
