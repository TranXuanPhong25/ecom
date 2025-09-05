package repositories

import (
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
)

func ConnectRedis(host string) {

	rdb = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
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
