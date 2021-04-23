package response

import (
	"github.com/dgb8901/go-wechat-miniapp-sdk/models"
)

type GetPaidUnionId struct {
	models.WxError
	Unionid string `json:"unionid"`
}
