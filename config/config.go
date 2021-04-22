package config

type WxaConfig interface {

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
}
