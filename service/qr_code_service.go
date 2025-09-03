package service

import (
	"github.com/dgb8901/go-wechat-miniapp-sdk/models/request"
)

const (
	// 获取小程序码
	createQrCodeUrl string = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode"
)

// WxaQrCodeService 小程序码
type WxaQrCodeService struct {
	wxaService *WxaService
}

func newWxaQrCodeService(service *WxaService) *WxaQrCodeService {
	return &WxaQrCodeService{wxaService: service}
}

// CreateQRCode 获取小程序二维码，适用于需要的码数量较少的业务场景。
// 通过该接口生成的小程序码，永久有效，有数量限制
func (s *WxaQrCodeService) CreateQRCode(param *request.CreateQrCode) ([]byte, error) {
	var params = map[string]interface{}{
		"path":  param.Path,
		"width": param.Width,
	}

	qrCode, err := s.wxaService.PostFile(createQrCodeUrl, &params)

	return qrCode, err
}
