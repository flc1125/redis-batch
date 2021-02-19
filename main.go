package main

import (
	"github.com/flc1125/redis-batch-delete/command"
)

func main() {
	command.Run()
}

// main -a -h 192.168.1.1 -p 12313 "*" --delete
