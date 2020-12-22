package main

import (
	"log"
	"os"

	"github.com/flc1125/redis-batch-delete/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "boom",
		Usage:  "make an explosive entrance",
		Action: commands.Demo,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
