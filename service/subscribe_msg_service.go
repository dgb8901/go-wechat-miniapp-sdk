package service

import (
	"errors"

	"github.com/dgb8901/go-wechat-miniapp-sdk/common"
	"github.com/dgb8901/go-wechat-miniapp-sdk/models"
	"github.com/dgb8901/go-wechat-miniapp-sdk/models/request"
)

const (
	// 订阅消息接口地址
	subscribeMessageUrl string = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"
)

type WxaSubscribeMsgService struct {
	wxaService *WxaService
}

func newWxaSubscribeMsgService(service *WxaService) *WxaSubscribeMsgService {
	return &WxaSubscribeMsgService{wxaService: service}
}

// Send 发送订阅消息
func (s *WxaSubscribeMsgService) Send(subMsg *request.SubscribeMsg) (*models.WxError, error) {
	var result models.WxError
	data := common.JsonToMap(common.ToJson(subMsg))
	err := s.wxaService.Post(subscribeMessageUrl, &data, &result)

	if err != nil {
		err = errors.New("Failed to send subscription message: " + err.Error())
		return nil, err
	}

	if result.ErrCode != 0 {
		err = errors.New(result.ErrMsg)
		return nil, err
	}

	return &result, nil
}
