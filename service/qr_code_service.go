package service

import (
	"go-wechat-miniapp-sdk/models/request"
)

const (
	// 获取小程序码
	createQrCodeUrl string = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode"
)

// 小程序码
type QrCodeService struct {
	wxaService *WxaService
}

// 获取小程序二维码，适用于需要的码数量较少的业务场景。
// 通过该接口生成的小程序码，永久有效，有数量限制
func (qr *QrCodeService) CreateQRCode(param *request.CreateQrCode) ([]byte, error) {
	var params = map[string]interface{}{
		"path":  param.Path,
		"width": param.Width,
	}

	qrCode, err := qr.wxaService.PostFile(createQrCodeUrl, &params)

	return qrCode, err
}
