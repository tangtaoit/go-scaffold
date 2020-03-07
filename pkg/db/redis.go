package db

import "github.com/tgo-team/FeiGeIMServer/pkg/redis"

// NewRedis 创建redis
func NewRedis(addr string) *redis.Conn {
	return redis.New(addr)
}
