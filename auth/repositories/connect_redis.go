package repositories

import "github.com/redis/go-redis/v9"

func ConnectRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "auth-redis:6379",
		Password: "redis-password", // no password set
		DB:       0,                // use default DB
	})
}
