package api

import (
	"github.com/gocraft/dbr/v2"
	"github.com/tgo-team/FeiGeIMServer/internal/common"
	"github.com/tgo-team/FeiGeIMServer/internal/config"
	"github.com/tgo-team/FeiGeIMServer/pkg/cache"
	"github.com/tgo-team/FeiGeIMServer/pkg/lmevent"
	"github.com/tgo-team/FeiGeIMServer/pkg/log"
	"github.com/tgo-team/FeiGeIMServer/pkg/pool"
	"github.com/tgo-team/FeiGeIMServer/pkg/redis"
)

// Context 配置上下文
type Context struct {
	cfg            *config.Config
	sqlLiteSession *dbr.Session
	redisCache     *common.RedisCache
	memoryCache    cache.Cache
	log.Log
	EventPool pool.Collector
	Event     lmevent.Event
}

// NewContext NewContext
func NewContext(cfg *config.Config) *Context {
	return &Context{cfg: cfg, Log: log.NewTLog("Context"), EventPool: pool.StartDispatcher(cfg.MessagePoolSize)}
}

// GetConfig 获取配置信息
func (c *Context) GetConfig() *config.Config {
	return c.cfg
}

// DB DB
func (c *Context) DB() *dbr.Session {
	return nil
}

// NewRedisCache 创建一个redis缓存
func (c *Context) NewRedisCache() *common.RedisCache {
	if c.redisCache == nil {
		c.redisCache = common.NewRedisCache(c.cfg.RedisAddr)
	}
	return c.redisCache
}

// NewMemoryCache 创建一个内存缓存
func (c *Context) NewMemoryCache() cache.Cache {
	if c.memoryCache == nil {
		c.memoryCache = common.NewMemoryCache()
	}
	return c.memoryCache
}

// Cache 缓存
func (c *Context) Cache() cache.Cache {
	if c.cfg.Test {
		return c.NewMemoryCache()
	}
	return c.NewRedisCache()
}

// GetRedisConn 获取redis连接
func (c *Context) GetRedisConn() *redis.Conn {
	return c.NewRedisCache().GetRedisConn()
}

// EventBegin 开启事件
func (c *Context) EventBegin(data *lmevent.Data, tx *dbr.Tx) (int64, error) {
	return c.Event.Begin(data, tx)
}

// EventCommit 提交事件
func (c *Context) EventCommit(eventID int64) {
	c.Event.Commit(eventID)
}
