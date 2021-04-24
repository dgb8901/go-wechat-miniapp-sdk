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

// 获取appid
func (config *wxaInMemoryConfig) GetAppId() string {
	return config.cfg.AppId
}

// 获取Secret
func (config *wxaInMemoryConfig) GetSecret() string {
	return config.cfg.Secret
}

// 获取access_token
func (config *wxaInMemoryConfig) GetAccessToken() string {
	return config.accessToken
}

// access_token是否过期
func (config *wxaInMemoryConfig) IsAccessTokenExpired() bool {
	return time.Now().Unix() > config.expiresTime
}

// 强制过期access_token
func (config *wxaInMemoryConfig) ExpiredAccessToken() {
	config.expiresTime = 0
}

// 更新access_token
func (config *wxaInMemoryConfig) UpdateAccessToken(accessToken string, expiresInSeconds int64) {
	config.accessToken = accessToken
	config.expiresTime = time.Now().Unix() + (expiresInSeconds - 200)
}

func (config *wxaInMemoryConfig) GetConfig() *Cfg {
	return config.cfg
}

func (config *wxaInMemoryConfig) SetConfig(cfg *Cfg) {
	config.cfg = cfg
}
