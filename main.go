package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/flc1125/redis-batch-delete/commands"
	"github.com/go-redis/redis/v8"
	"github.com/urfave/cli/v2"
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

	redisClient()

	app := &cli.App{
		Name:   "boom",
		Usage:  "make an explosive entrance",
		Action: commands.Demo,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
