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

// 初始化redis配置
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

// 获取appid
func (config *wxaInRedisConfig) GetAppId() string {
	return config.cfg.AppId
}

// 获取Secret
func (config *wxaInRedisConfig) GetSecret() string {
	return config.cfg.Secret
}

// 获取access_token
func (config *wxaInRedisConfig) GetAccessToken() string {
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
func (config *wxaInRedisConfig) IsAccessTokenExpired() bool {
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
func (config *wxaInRedisConfig) ExpiredAccessToken() {
	conn := config.pool.Get()
	defer conn.Close()
	do, err := conn.Do("DEL", config.cfg.AppId)
	log.Printf("redis-->Delete:%s", do)
	if err != nil {
		log.Printf("redis-->Delete:%s", err.Error())
	}
}

// 更新access_token
func (config *wxaInRedisConfig) UpdateAccessToken(accessToken string, expiresInSeconds int64) {
	conn := config.pool.Get()
	defer conn.Close()
	expireTime := time.Now().Unix() + (expiresInSeconds - 200)
	do, err := conn.Do("SET", config.cfg.AppId, accessToken, "EX", expireTime)
	log.Printf("redis-->SetExpire:%s", do)
	if err != nil {
		log.Printf("redis-->SetExpire:%s", err.Error())
	}
}

func (config *wxaInRedisConfig) GetConfig() *Cfg {
	return config.cfg
}

func (config *wxaInRedisConfig) SetConfig(cfg *Cfg) {
	config.cfg = cfg
}
