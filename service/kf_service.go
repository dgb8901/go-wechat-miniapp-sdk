package service

import (
	"errors"
	"github.com/dgb8901/go-wechat-miniapp-sdk/common"
	"github.com/dgb8901/go-wechat-miniapp-sdk/models"
	"github.com/dgb8901/go-wechat-miniapp-sdk/models/request"
	"log"
)

const (
	getTempMedia string = "https://api.weixin.qq.com/cgi-bin/media/get"
	sendKfMsg    string = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
	setTyping    string = "https://api.weixin.qq.com/cgi-bin/message/custom/typing"
)

// 客服
type wxaKfService struct {
	wxaService *WxaService
}

// GetTempMedia 获取客服消息内的临时素材,即下载临时的多媒体文件。目前小程序仅支持下载图片文件。
func (kf *wxaKfService) GetTempMedia(mediaId string) ([]byte, error) {

	if common.IsBlank(mediaId) {
		return nil, errors.New("mediaId is blank")
	}
	var params = map[string]interface{}{
		"media_id": mediaId,
	}

	resp, err := kf.wxaService.GetFile(getTempMedia, &params)

	return resp, err
}

// Send 发送客服消息
func (kf *wxaKfService) Send(message *request.SendKfMessage) (*models.WxError, error) {

	var result models.WxError
	data := common.JsonToMap(common.ToJson(message))

	err := kf.wxaService.Post(sendKfMsg, &data, &result)

	if err != nil {
		err = errors.New("failed to send kf message：" + err.Error())
		log.Printf("failed to send kf message：%s", err.Error())
		return nil, err
	}

	if result.ErrCode != 0 {
		err = errors.New(result.ErrMsg)
		log.Printf(result.ErrMsg)
		return nil, err
	}

	return &result, nil
}

// SetTyping 下发客服当前输入状态给用户
func (kf *wxaKfService) SetTyping(typing request.SetTyping) (*models.WxError, error) {

	if typing.Command != "Typing" {
		return nil, errors.New("command invalid")
	} else if typing.Command != "CancelTyping" {
		return nil, errors.New("command invalid")
	}

	var result models.WxError
	data := common.JsonToMap(common.ToJson(typing))

	err := kf.wxaService.Post(setTyping, &data, &result)

	if err != nil {
		err = errors.New("failed to send kf message status：" + err.Error())
		log.Printf("failed to send kf message status：%s", err.Error())
		return nil, err
	}

	if result.ErrCode != 0 {
		err = errors.New(result.ErrMsg)
		log.Printf(result.ErrMsg)
		return nil, err
	}

	return &result, nil
}
