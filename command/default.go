package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// 默认文案
func Default(c *cli.Context) error {
	fmt.Println("请输入XXX")

	return nil
}
