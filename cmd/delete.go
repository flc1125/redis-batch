package cmd

import (
	"context"
	"fmt"
	"github.com/flc1125/redis-batch-delete/pkg/redis"
	"github.com/spf13/cobra"
	"strconv"
	"sync"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "批量删除 Redis Key",
	Run: func(cmd *cobra.Command, args []string) {
		logic := &deleteLogic{
			keys: make(chan []string, 1000),
			wg:   sync.WaitGroup{},
		}

		logic.delete()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var ctx = context.Background()

type deleteLogic struct {
	keys chan []string
	wg   sync.WaitGroup
}

func (this *deleteLogic) delete() {
	defer this.wg.Wait()

	go this.del()

	this.scan()
}

func (this *deleteLogic) del() {
	for {
		keys := <-this.keys

		val, err := redis.RDB.Del(ctx, keys...).Result()

		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprint(keys) + "::删除成功，删除数量:" + strconv.Itoa(int(val)))

		this.wg.Done()
	}
}

func (this *deleteLogic) scan() {
	var (
		keys   []string
		cursor uint64
		err    error
	)

	for {
		keys, cursor, err = redis.RDB.Scan(ctx, cursor, "*", 20).Result()

		if err != nil {
			panic(err)
		}

		if len(keys) == 0 {
			break
		}

		this.wg.Add(1)
		this.keys <- keys

		if cursor == 0 {
			break
		}
	}
}