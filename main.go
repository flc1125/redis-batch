package main

import (
	"context"
	"fmt"

	"github.com/flc1125/redis-batch-delete/command"
	"github.com/flc1125/redis-batch-delete/config"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func redisClient() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.8.117:7003"},
	})

	var keys []string
	var cursor uint64
	var count int

	for {
		keys, cursor, _ = rdb.Scan(ctx, cursor, "*2019*", 100).Result()

		if len(keys) == 0 {
			break
		}

		_, err := rdb.Del(ctx, keys...).Result()

		fmt.Println(keys, err)

		count += len(keys)
		break
	}

	fmt.Println(count)
}

func main() {

	// redisClient()
	command.Run(&config.Config{
		Host:    "localhost",
		Port:    6379,
		Cluster: false,
	})
}

// main -a -h 192.168.1.1 -p 12313 "*" --delete
