package config

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

// 基于redis配置
type wxaInRedisConfig struct {
	cfg  *Cfg
	pool *redis.Pool
}

// NewInRedis 初始化redis配置
func NewInRedis(cfg *Cfg, server, password string) *wxaInRedisConfig {
	pool := &redis.Pool{
		MaxIdle:     5, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   10, //最大数
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server,
				redis.DialReadTimeout(time.Second*10),
				redis.DialConnectTimeout(time.Second*30),
				redis.DialPassword(password),
				redis.DialDatabase(0),
			)
			if err != nil {
				panic(err)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")

			if err != nil {
				log.Printf("testOnBorrow:%s", err.Error())
			}
			return err
		},
	}
	return &wxaInRedisConfig{cfg: cfg, pool: pool}
}

// GetAppId 获取appid
func (c *wxaInRedisConfig) GetAppId() string {
	return c.cfg.AppId
}

// GetSecret 获取Secret
func (c *wxaInRedisConfig) GetSecret() string {
	return c.cfg.Secret
}

// GetAccessToken 获取access_token
func (c *wxaInRedisConfig) GetAccessToken() string {
	conn := c.pool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", c.cfg.AppId))
	log.Printf("redis-->GetValue:%s", value)
	if err != nil {
		value = ""
		log.Printf("err redis-->Get:%s", err.Error())
	}
	return value
}

// IsAccessTokenExpired access_token是否过期
func (c *wxaInRedisConfig) IsAccessTokenExpired() bool {
	conn := c.pool.Get()
	defer conn.Close()
	b, err := redis.Bool(conn.Do("EXISTS", c.cfg.AppId))
	if err != nil {
		log.Printf("redis-->Exists:%s", err.Error())
		return false
	}
	return b
}

// ExpiredAccessToken 强制过期access_token
func (c *wxaInRedisConfig) ExpiredAccessToken() {
	conn := c.pool.Get()
	defer conn.Close()
	do, err := conn.Do("DEL", c.cfg.AppId)
	log.Printf("redis-->Delete:%s", do)
	if err != nil {
		log.Printf("redis-->Delete:%s", err.Error())
	}
}

// UpdateAccessToken 更新access_token
func (c *wxaInRedisConfig) UpdateAccessToken(accessToken string, expiresInSeconds int64) {
	conn := c.pool.Get()
	defer conn.Close()
	expireTime := time.Now().Unix() + (expiresInSeconds - 200)
	do, err := conn.Do("SET", c.cfg.AppId, accessToken, "EX", expireTime)
	log.Printf("redis-->SetExpire:%s", do)
	if err != nil {
		log.Printf("redis-->SetExpire:%s", err.Error())
	}
}

func (c *wxaInRedisConfig) GetConfig() *Cfg {
	return c.cfg
}

func (c *wxaInRedisConfig) SetConfig(cfg *Cfg) {
	c.cfg = cfg
}
