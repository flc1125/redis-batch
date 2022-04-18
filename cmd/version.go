package cmd

import (
	"fmt"
	"github.com/flc1125/redis-batch/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Redis-Batch v" + version.VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
