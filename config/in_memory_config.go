package config

import "time"

// 基于内存配置
type wxaInMemoryConfig struct {
	cfg         *Cfg
	accessToken string
	expiresTime int64
}

func NewInMemory(cfg *Cfg) *wxaInMemoryConfig {
	return &wxaInMemoryConfig{cfg: cfg}
}

// GetAppId 获取appid
func (c *wxaInMemoryConfig) GetAppId() string {
	return c.cfg.AppId
}

// GetSecret 获取Secret
func (c *wxaInMemoryConfig) GetSecret() string {
	return c.cfg.Secret
}

// GetAccessToken 获取access_token
func (c *wxaInMemoryConfig) GetAccessToken() string {
	return c.accessToken
}

// IsAccessTokenExpired access_token是否过期
func (c *wxaInMemoryConfig) IsAccessTokenExpired() bool {
	return time.Now().Unix() > c.expiresTime
}

// ExpiredAccessToken 强制过期access_token
func (c *wxaInMemoryConfig) ExpiredAccessToken() {
	c.expiresTime = 0
}

// UpdateAccessToken 更新access_token
func (c *wxaInMemoryConfig) UpdateAccessToken(accessToken string, expiresInSeconds int64) {
	c.accessToken = accessToken
	c.expiresTime = time.Now().Unix() + (expiresInSeconds - 200)
}

func (c *wxaInMemoryConfig) GetConfig() *Cfg {
	return c.cfg
}

func (c *wxaInMemoryConfig) SetConfig(cfg *Cfg) {
	c.cfg = cfg
}
