package command

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// worker 任务处理
func worker(rdb *redis.Client, queue chan string) {
	for {
		key := <-queue

		_, err := rdb.Del(ctx, key).Result()

		if err != nil {
			fmt.Println("failed to delete:", key)
			continue
		}

		fmt.Println("Deleted：", key)
	}
}

// creater 创建任务者
func creater(rdb *redis.Client, queue chan string) {
	var keys []string
	var cursor uint64 = 0
	var err error

	for {
		keys, cursor, err = rdb.Scan(ctx, cursor, "*", 5).Result()

		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			queue <- key
			fmt.Println("Found：", key)
		}

		if cursor == 0 {
			break
		}
	}
}

// Handle 处理查询与清理
func Handle(rdb *redis.Client) {
	var queue = make(chan string)

	go worker(rdb, queue)

	go creater(rdb, queue)

	time.Sleep(time.Second * 5)
}
