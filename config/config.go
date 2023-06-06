package config

type Cfg struct {
	AppId         string
	Secret        string
	Token         string
	AesKey        string
	MsgDataFormat string
}

type Config interface {

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
	GetConfig() *Cfg
	// SetConfig 设置配置信息
	SetConfig(cfg *Cfg)
}
