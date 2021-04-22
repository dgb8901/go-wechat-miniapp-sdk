package response

import "go-wechat-miniapp-sdk/models"

type JsCode2SessionResult struct {
	models.WxError
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	ExpiresIn  int64  `json:"expires_in"`
	SessionKey string `json:"session_key"`
}
