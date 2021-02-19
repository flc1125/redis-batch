package command

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// 创建清理的协程
func worker(rdb *redis.Client, queue chan string) {
	for {
		key := <-queue

		fmt.Println(key, "Deleted")
		_, err := rdb.Del(ctx, key).Result()

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func Handle(rdb *redis.Client) {
	var queue = make(chan string)
	var keys []string
	var cursor uint64 = 0
	var err error

	go worker(rdb, queue)

	for {
		keys, cursor, err = rdb.Scan(ctx, cursor, "*", 5).Result()

		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			// queue = append(queue, key)
			queue <- key
		}

		if cursor == 0 {
			break
		}
	}

	// fmt.Println(queue)
	time.Sleep(time.Second * 5)
}
