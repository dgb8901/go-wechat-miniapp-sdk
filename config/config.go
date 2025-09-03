package config

type Config struct {
	AppId         string
	Secret        string
	Token         string
	AesKey        string
	MsgDataFormat string

	// Server Redis host 地址 如：127.0.0.1:6379
	Server string
	// Password Redis 密码
	Password string
}

type CfgInterface interface {

	// GetAppId 获取appid
	GetAppId() string

	// GetSecret 获取Secret
	GetSecret() string

	// GetAccessToken 获取access_token
	GetAccessToken() string

	// IsAccessTokenExpired access_token是否过期
	IsAccessTokenExpired() bool

	// ExpiredAccessToken 强制过期access_token
	ExpiredAccessToken()

	// UpdateAccessToken 更新access_token
	UpdateAccessToken(accessToken string, expiresInSeconds int64)

	// GetConfig 获取配置信息
	GetConfig() *Config
	// SetConfig 设置配置信息
	SetConfig(cfg *Config)
}
