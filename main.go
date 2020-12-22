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
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	keys, _, _ := rdb.Scan(ctx, 0, "*", 10).Result()

	fmt.Println(keys)
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
