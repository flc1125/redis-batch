package cmd

import (
	"fmt"
	"github.com/flc1125/redis-batch/pkg/redis"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var generateTestsCmd = &cobra.Command{
	Use:   "generate:tests",
	Short: "批量生成测试数据",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 1; i <= 10000; i++ {
			redis.RDB.Set(ctx, "tests:"+strconv.Itoa(i), "value:"+strconv.Itoa(i), 5*time.Minute).Result()

			fmt.Println(i)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateTestsCmd)
}
