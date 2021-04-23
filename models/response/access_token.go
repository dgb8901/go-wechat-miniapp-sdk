package response

import (
	"github.com/dgb8901/go-wechat-miniapp-sdk/models"
)

type WxaAccessToken struct {
	models.WxError
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
