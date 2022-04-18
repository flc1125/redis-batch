package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "redis-batch",
	Short: "redis-batch 命令",
	Long:  `此处省略 80000 字介绍`,
}

var ctx = context.Background()
