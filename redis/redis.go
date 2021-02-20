package redis

import "github.com/go-redis/redis/v8"

// New 创建 redis 连接
func New() *redis.Client {
	return NewClient()
}

// NewClient 创建单机版连接
func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "192.168.2.126:6379",
		Password: "",
		DB:       0,
	})
}
