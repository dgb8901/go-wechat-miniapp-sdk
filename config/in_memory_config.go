package config

import "time"

// MemoryConfig 基于内存配置
type MemoryConfig struct {
	cfg         *Config
	accessToken string
	expiresTime int64
}

func NewInMemory(cfg *Config) CfgInterface {
	return &MemoryConfig{cfg: cfg}
}

// GetAppId 获取appid
func (c *MemoryConfig) GetAppId() string {
	return c.cfg.AppId
}

// GetSecret 获取Secret
func (c *MemoryConfig) GetSecret() string {
	return c.cfg.Secret
}

// GetAccessToken 获取access_token
func (c *MemoryConfig) GetAccessToken() string {
	return c.accessToken
}

// IsAccessTokenExpired access_token是否过期
func (c *MemoryConfig) IsAccessTokenExpired() bool {
	return time.Now().Unix() > c.expiresTime
}

// ExpiredAccessToken 强制过期access_token
func (c *MemoryConfig) ExpiredAccessToken() {
	c.expiresTime = 0
}

// UpdateAccessToken 更新access_token
func (c *MemoryConfig) UpdateAccessToken(accessToken string, expiresInSeconds int64) {
	c.accessToken = accessToken
	c.expiresTime = time.Now().Unix() + (expiresInSeconds - 200)
}

func (c *MemoryConfig) GetConfig() *Config {
	return c.cfg
}

func (c *MemoryConfig) SetConfig(cfg *Config) {
	c.cfg = cfg
}
