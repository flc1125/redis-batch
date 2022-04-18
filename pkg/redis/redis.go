package redis

import (
	redisV8 "github.com/go-redis/redis/v8"
)

var RDB *redisV8.Client

func init() {
	RDB = redisV8.NewClient(&redisV8.Options{
		Addr:     "localhost:63790",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
