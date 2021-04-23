package config

type Config struct {
	AppId         string
	Secret        string
	Token         string
	AesKey        string
	MsgDataFormat string
	AccessToken   string
	ExpiresTime   int64
}

type configInterface interface {

	// 获取appid
	GetAppId() string

	// 获取Secret
	GetSecret() string

	// 获取access_token
	GetAccessToken() string

	// access_token是否过期
	IsAccessTokenExpired() bool

	// 强制过期access_token
	ExpiredAccessToken()

	// 更新access_token
	UpdateAccessToken(accessToken string, expiresInSeconds int64)

	// 获取配置信息
	GetConfig() *Config
	// 设置配置信息
	SetConfig(cfg *Config)
}
