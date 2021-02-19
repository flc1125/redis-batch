package redis

import "github.com/go-redis/redis/v8"

// 创建 redis 连接
func New() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
