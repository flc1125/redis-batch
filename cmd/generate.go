package cmd

import (
	"fmt"
	"github.com/flc1125/redis-batch/pkg/redis"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "生成测试数据；如果你的服务器没有足够大的内存，请艾特你的运维大大，并暗示他赶紧买。",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 1; i <= 10000; i++ {
			redis.RDB.Set(ctx, "tests:"+strconv.Itoa(i), "value:"+strconv.Itoa(i), 5*time.Minute).Result()

			fmt.Println(i)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
