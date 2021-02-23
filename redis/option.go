package redis

import "github.com/go-redis/redis/v8"

type Option struct {
	Type           string
	Cluster        bool
	Options        *redis.Options
	ClusterOptions *redis.ClusterOptions
}

// func (o *Option) GetOptions() *redis.Options {
// 	return o.Options
// }

// func (o *Option) GetClusterOptions() *redis.ClusterOptions {
// 	return o.ClusterOptions
// }
