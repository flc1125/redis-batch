package cmd

import (
	"fmt"
	"github.com/flc1125/redis-batch/pkg/redis"
	"github.com/flc1125/redis-batch/util/coroutine"
	"github.com/spf13/cobra"
	"strconv"
	"sync"
	"time"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "生成测试数据；如果你的服务器没有足够大的内存，请艾特你的运维大大，并暗示他赶紧买。",
	Run: func(cmd *cobra.Command, args []string) {
		co := coroutine.New(20)
		wg := sync.WaitGroup{}
		defer wg.Wait()

		for i := 1; i <= 10000; i++ {
			wg.Add(1)

			ii := i // 避免出现变量循环

			co.Run(func(co *coroutine.Coroutine) {
				defer wg.Done()

				redis.RDB.Set(ctx, "tests:"+strconv.Itoa(ii), "value:"+strconv.Itoa(ii), 5*time.Minute).Result()

				fmt.Println(ii)
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
