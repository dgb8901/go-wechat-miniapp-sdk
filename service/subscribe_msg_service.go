package service

import (
	"errors"
	"go-wechat-miniapp-sdk/common"
	"go-wechat-miniapp-sdk/models"
	"go-wechat-miniapp-sdk/models/request"
	"log"
)

const (
	// 订阅消息接口地址
	subscribeMessageUrl string = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"
)

type WxaSubscribeMsgService struct {
	wxaService *WxaService
}

// 发送订阅消息
func (msg *WxaSubscribeMsgService) Send(subMsg *request.SubscribeMsg) (*models.WxError, error) {
	var result models.WxError
	data := common.JsonToMap(common.ToJson(subMsg))
	err := msg.wxaService.Post(subscribeMessageUrl, &data, &result)

	if err != nil {
		err = errors.New("Failed to send subscription message：" + err.Error())
		log.Printf("Failed to send subscription message：%s", err.Error())
		return nil, err
	}

	if result.ErrCode != 0 {
		err = errors.New(result.ErrMsg)
		log.Printf(result.ErrMsg)
		return nil, err
	}

	return &result, nil
}
