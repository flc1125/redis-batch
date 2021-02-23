package redis

import "github.com/go-redis/redis/v8"

type Redis struct {
	Option        *Option
	Client        *redis.Client
	ClusterClient *redis.ClusterClient
}

func New(o *Option) *Redis {
	r := &Redis{
		Option: o,
	}

	switch o.Cluster {
	case true:
		r.ClusterClient = redis.NewClusterClient(o.ClusterOptions)
	default:
		r.Client = redis.NewClient(o.Options)
	}

	return r
}

func NewClient(o *Option) *Redis.Client {
	redis := New(o)

	return redis.Client
}
