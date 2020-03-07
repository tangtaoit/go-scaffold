package config

import (
	"os"
	"strings"
	"time"
)

// Config 配置信息
type Config struct {
	TokenExpire      time.Duration // token失效时间
	SQLDir           string        // 数据库脚本路径
	RedisAddr        string        // redis地址
	Test             bool          // 是否是测试模式
	IMUrl            string        // im基地址
	TokenCachePrefix string        // token缓存前缀
	MessagePoolSize  int64         // 发消息任务池大小
	APNSDev          bool          // apns是否是开发模式
	APNSPassword     string        // apns的密码
	APNSTopic        string
	NameCacheExpire  time.Duration // 名字缓存过期时间
}

// New New
func New() *Config {
	return &Config{
		TokenExpire:      time.Hour * 24 * 30,
		SQLDir:           "configs/sql",
		RedisAddr:        GetEnv("REDIS_ADDR", "127.0.0.1:6379"),
		Test:             false,
		IMUrl:            GetEnv("IM_URL", "http://127.0.0.1:8029"),
		TokenCachePrefix: "token:",
		MessagePoolSize:  100,
		APNSDev:          true,
		APNSPassword:     "123456",
		APNSTopic:        "com.limao.im.LiMaoIMDemo",
		NameCacheExpire:  time.Hour * 24 * 7,
	}
}

// GetEnv 成环境变量里获取
func GetEnv(key string, defaultValue string) string {
	v := os.Getenv(key)
	if strings.TrimSpace(v) == "" {
		return defaultValue
	}
	return v
}
