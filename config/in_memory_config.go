package config

import "time"

// 基于内存配置
type WxaInMemoryConfig struct {
	cfg *Config
}

func NewMemory(cfg *Config) *WxaInMemoryConfig {
	return &WxaInMemoryConfig{cfg: cfg}
}

// 获取appid
func (config *WxaInMemoryConfig) GetAppId() string {
	return config.cfg.AppId
}

// 获取Secret
func (config *WxaInMemoryConfig) GetSecret() string {
	return config.cfg.Secret
}

// 获取access_token
func (config *WxaInMemoryConfig) GetAccessToken() string {
	return config.cfg.AccessToken
}

// access_token是否过期
func (config *WxaInMemoryConfig) IsAccessTokenExpired() bool {
	return time.Now().Unix() > config.cfg.ExpiresTime
}

// 强制过期access_token
func (config *WxaInMemoryConfig) ExpiredAccessToken() {
	config.cfg.ExpiresTime = 0
}

// 更新access_token
func (config *WxaInMemoryConfig) UpdateAccessToken(accessToken string, expiresInSeconds int64) {
	config.cfg.AccessToken = accessToken
	config.cfg.ExpiresTime = time.Now().Unix() + (expiresInSeconds - 200)
}

func (config *WxaInMemoryConfig) GetConfig() *Config {
	return config.cfg
}

func (config *WxaInMemoryConfig) SetConfig(cfg *Config) {
	config.cfg = cfg
}
