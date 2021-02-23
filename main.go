package main

import "github.com/flc1125/redis-batch-delete/redis"

func main() {
	redis.New(&redis.Option{})
	// command.Run()
}

// main -a -h 192.168.1.1 -p 12313 "*" --delete
