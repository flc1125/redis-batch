package main

import "github.com/go-redis/redis/v8"

func createClusterClient() *redis.ClusterClient {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.8.117:7003"},
	})

	return rdb
}

func RedisClient() {

}
