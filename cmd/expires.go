package cmd

import (
	"fmt"
	"github.com/flc1125/redis-batch/pkg/redis"
	"github.com/flc1125/redis-batch/util/coroutine"
	"github.com/spf13/cobra"
	"sync"
	"time"
)

var expiresCmd = &cobra.Command{
	Use:   "expires",
	Short: "批量设置过期",
	Run: func(cmd *cobra.Command, args []string) {
		logic := &expiresLogic{
			key: make(chan string, 1000),
			wg:  sync.WaitGroup{},
			co:  coroutine.New(20),
		}

		logic.delete()
	},
}

func init() {
	rootCmd.AddCommand(expiresCmd)
}

type expiresLogic struct {
	key chan string
	wg  sync.WaitGroup
	co  *coroutine.Coroutine
}

func (this *expiresLogic) delete() {
	defer this.wg.Wait()

	go this.expires()

	this.scan()
}

func (this *expiresLogic) expires() {
	for {
		key := <-this.key

		this.co.Run(func(co *coroutine.Coroutine) {
			defer this.wg.Done()

			_, err := redis.RDB.Expire(ctx, key, time.Minute*5).Result()

			if err != nil {
				return
			}

			fmt.Println(key + "::设置过期时间成功")
		})
	}
}

func (this *expiresLogic) scan() {
	var (
		keys   []string
		cursor uint64
		err    error
	)

	for {
		keys, cursor, err = redis.RDB.Scan(ctx, cursor, "*", 100).Result()

		if err != nil {
			panic(err)
		}

		if len(keys) == 0 {
			break
		}

		for _, key := range keys {
			this.wg.Add(1)
			this.key <- key
		}

		if cursor == 0 {
			break
		}
	}
}
