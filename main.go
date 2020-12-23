package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

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

	// redisClient()

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "localhost",
				Usage: "HOST 信息",
			},
			&cli.IntFlag{
				Name:    "port",
				Value:   6379,
				Aliases: []string{"p"},
				Usage:   "PORT 信息",
			},
			&cli.BoolFlag{
				Name:    "cluster",
				Value:   false,
				Aliases: []string{"c"},
				Usage:   "Cluster 集群模式",
			},
			&cli.BoolFlag{
				Name:    "show",
				Value:   false,
				Aliases: []string{"s"},
				Usage:   "查看模式",
			},
		},
		Name:  "redis-batch-delete",
		Usage: "Redis 批量删除操作",
		Action: func(c *cli.Context) error {
			host, port, cluster := c.String("host"), c.Int("port"), c.Bool("cluster")

			fmt.Println("afasdfasf---" + host + ":" + strconv.Itoa(port) + "-" + strconv.FormatBool(cluster))

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// main -a -h 192.168.1.1 -p 12313 "*" --delete
