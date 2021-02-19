package command

import (
	"log"
	"os"

	"github.com/flc1125/redis-batch-delete/redis"
	"github.com/urfave/cli/v2"
)

func Run() {
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
			rdb := redis.New()

			Handle(rdb)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
