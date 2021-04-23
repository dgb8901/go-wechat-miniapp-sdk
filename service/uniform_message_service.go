package service

import (
	"errors"
	"github.com/dgb8901/go-wechat-miniapp-sdk/common"
	"github.com/dgb8901/go-wechat-miniapp-sdk/models"
	"github.com/dgb8901/go-wechat-miniapp-sdk/models/request"
	"log"
)

const (
	// 订阅消息接口地址
	uniformSendUrl string = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send"
)

// 统一服务消息
type wxaUniformMessageService struct {
	wxaService *WxaService
}

// 发送服务消息
func (un *wxaUniformMessageService) Send(message *request.UniformMessage) (*models.WxError, error) {
	var result models.WxError
	data := common.JsonToMap(common.ToJson(message))
	err := un.wxaService.Post(uniformSendUrl, &data, &result)

	if err != nil {
		err = errors.New("Failed to send service message：" + err.Error())
		log.Printf("Failed to send service message：%s", err.Error())
		return nil, err
	}

	if result.ErrCode != 0 {
		err = errors.New(result.ErrMsg)
		log.Printf(result.ErrMsg)
		return nil, err
	}

	return &result, nil
}
