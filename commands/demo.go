package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Demo(c *cli.Context) error {
	fmt.Println("Hello world")
	return nil
}
