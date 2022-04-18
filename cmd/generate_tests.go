package cmd

import (
	"fmt"
	"github.com/flc1125/redis-batch-delete/pkg/redis"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var generateTestsCmd = &cobra.Command{
	Use:   "generate:tests",
	Short: "批量生成测试数据",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 1; i <= 10000; i++ {
			redis.RDB.Set(ctx, "tests:"+strconv.Itoa(i), "value:"+strconv.Itoa(i), time.Minute*5).Result()

			fmt.Println(i)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateTestsCmd)
}
