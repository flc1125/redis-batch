package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Value:   "english",
				Aliases: []string{"l", "ln"},
				Usage:   "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {

			fmt.Println("Hello friend!" + c.String("lang"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
