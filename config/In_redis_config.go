package config

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

// 基于redis配置
type WxaInRedisConfig struct {
	cfg  *Config
	pool *redis.Pool
}

func NewRedis(cfg *Config, server, password string) *WxaInRedisConfig {
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
	return &WxaInRedisConfig{cfg: cfg, pool: pool}
}

// 获取appid
func (config *WxaInRedisConfig) GetAppId() string {
	return config.cfg.AppId
}

// 获取Secret
func (config *WxaInRedisConfig) GetSecret() string {
	return config.cfg.Secret
}

// 获取access_token
func (config *WxaInRedisConfig) GetAccessToken() string {
	conn := config.pool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", config.cfg.AppId))
	log.Printf("redis-->GetValue:%s", value)
	if err != nil {
		value = ""
		log.Printf("err redis-->Get:%s", err.Error())
	}
	return value
}

// access_token是否过期
func (config *WxaInRedisConfig) IsAccessTokenExpired() bool {
	conn := config.pool.Get()
	defer conn.Close()
	b, err := redis.Bool(conn.Do("EXISTS", config.cfg.AppId))
	if err != nil {
		log.Printf("redis-->Exists:%s", err.Error())
		return false
	}
	return b
}

// 强制过期access_token
func (config *WxaInRedisConfig) ExpiredAccessToken() {
	conn := config.pool.Get()
	defer conn.Close()
	do, err := conn.Do("DEL", config.cfg.AppId)
	log.Printf("redis-->Delete:%s", do)
	if err != nil {
		log.Printf("redis-->Delete:%s", err.Error())
	}
}

// 更新access_token
func (config *WxaInRedisConfig) UpdateAccessToken(accessToken string, expiresInSeconds int64) {
	conn := config.pool.Get()
	defer conn.Close()
	expireTime := time.Now().Unix() + (expiresInSeconds - 200)
	do, err := conn.Do("SET", config.cfg.AppId, accessToken, "EX", expireTime)
	log.Printf("redis-->SetExpire:%s", do)
	if err != nil {
		log.Printf("redis-->SetExpire:%s", err.Error())
	}
}

func (config *WxaInRedisConfig) GetConfig() *Config {
	return config.cfg
}

func (config *WxaInRedisConfig) SetConfig(cfg *Config) {
	config.cfg = cfg
}
