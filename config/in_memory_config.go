package config

import "time"

// 基于内存配置
type WxaInMemoryConfig struct {
	AppId         string
	Secret        string
	Token         string
	AesKey        string
	MsgDataFormat string
	AccessToken   string
	ExpiresTime   int64
}

// 获取appid
func (config *WxaInMemoryConfig) GetAppId() string {
	return config.AppId
}

// 获取Secret
func (config *WxaInMemoryConfig) GetSecret() string {
	return config.Secret
}

// 获取access_token
func (config *WxaInMemoryConfig) GetAccessToken() string {
	return config.AccessToken
}

// access_token是否过期
func (config *WxaInMemoryConfig) IsAccessTokenExpired() bool {
	return time.Now().Unix() > config.ExpiresTime
}

// 强制过期access_token
func (config *WxaInMemoryConfig) ExpiredAccessToken() {
	config.ExpiresTime = 0
}

// 更新access_token
func (config *WxaInMemoryConfig) UpdateAccessToken(accessToken string, expiresInSeconds int64) {
	config.AccessToken = accessToken
	config.ExpiresTime = time.Now().Unix() + (expiresInSeconds - 200)
}
