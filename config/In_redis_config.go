package config

import "time"

// TODO 基于redis配置
type WxaInRedisConfig struct {
	cfg *Config
}

func NewRedis(cfg *Config) *WxaInRedisConfig {
	return &WxaInRedisConfig{cfg: cfg}
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
	return config.cfg.AccessToken
}

// access_token是否过期
func (config *WxaInRedisConfig) IsAccessTokenExpired() bool {
	return time.Now().Unix() > config.cfg.ExpiresTime
}

// 强制过期access_token
func (config *WxaInRedisConfig) ExpiredAccessToken() {
	config.cfg.ExpiresTime = 0
}

// 更新access_token
func (config *WxaInRedisConfig) UpdateAccessToken(accessToken string, expiresInSeconds int64) {
	config.cfg.AccessToken = accessToken
	config.cfg.ExpiresTime = time.Now().Unix() + (expiresInSeconds - 200)
}

func (config *WxaInRedisConfig) GetConfig() *Config {
	return config.cfg
}

func (config *WxaInRedisConfig) SetConfig(cfg *Config) {
	config.cfg = cfg
}
