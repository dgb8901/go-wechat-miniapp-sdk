package service

import (
	"errors"
	"log"

	"github.com/dgb8901/go-wechat-miniapp-sdk/common"
	"github.com/dgb8901/go-wechat-miniapp-sdk/models"
	"github.com/dgb8901/go-wechat-miniapp-sdk/models/request"
)

const (
	// 订阅消息接口地址
	uniformSendUrl string = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send"
)

// WxaUniformMessageService 统一服务消息
type WxaUniformMessageService struct {
	wxaService *WxaService
}

func newWxaUniformMessageService(service *WxaService) *WxaUniformMessageService {
	return &WxaUniformMessageService{wxaService: service}
}

// Send 发送服务消息
func (s *WxaUniformMessageService) Send(message *request.UniformMessage) (*models.WxError, error) {
	var result models.WxError
	data := common.JsonToMap(common.ToJson(message))
	err := s.wxaService.Post(uniformSendUrl, &data, &result)

	if err != nil {
		err = errors.New("Failed to send service message: " + err.Error())
		return nil, err
	}

	if result.ErrCode != 0 {
		err = errors.New(result.ErrMsg)
		log.Printf(result.ErrMsg)
		return nil, err
	}

	return &result, nil
}
