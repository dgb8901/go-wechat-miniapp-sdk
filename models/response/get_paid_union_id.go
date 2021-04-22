package response

import "go-wechat-miniapp-sdk/models"

type GetPaidUnionId struct {
	models.WxError
	Unionid string `json:"unionid"`
}
