package repositories

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"os"
)

var (
	rdb *redis.Client
)

func ConnectRedis() {

	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // no password set
		DB:       0,             // use default DB
	})
}

func TestRedis() {
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Print("Error connecting to redis\n")
	} else {
		log.Print("Connected to redis\n")
	}

}

func CloseRedisConnection() {
	if rdb != nil {
		if err := rdb.Close(); err != nil {
			log.Errorf("Error closing redis connection: %v", err)
		} else {
			log.Info("Redis connection closed successfully")
		}
	}
}
