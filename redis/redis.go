package redis

// import (
// 	"github.com/go-redis/redis/v8"
// )

// type Client interface{}

// type Options interface{}

// type Config struct {
// 	Cluster bool
// 	Options *Options
// }

// type Redis struct {
// 	config *Config
// 	client *Client
// }

// // New 创建 redis 连接
// func New(config *Config) *Client {
// 	var redis Redis

// 	redis.config = config

// 	switch config.Cluster {
// 	case true:
// 		redis.client = NewClusterClient(config.Options)
// 	default:
// 		redis.client = NewClient(config.Options)
// 	}

// 	return redis.client
// }

// // NewClient 创建单机版连接
// func NewClient(options *Options) *Client {
// 	opt := (*options).(redis.Options)
// 	return redis.NewClient(opt)
// }

// func NewClusterClient(options *Options) *Client {
// 	return redis.NewClusterClient(options)
// }
