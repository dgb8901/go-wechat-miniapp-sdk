package service

import (
	"github.com/pkg/errors"
	"go-wechat-miniapp-sdk/common"
	"go-wechat-miniapp-sdk/models/response"
	"log"
)

const (
	// 通过jscode获取openId
	jsocdeToSessionUrl string = "https://api.weixin.qq.com/sns/jscode2session"
	getPaidUnionId     string = "https://api.weixin.qq.com/wxa/getpaidunionid"
)

type wxaUserService struct {
	wxaService *wxaService
}

// 用户登录
func (user *wxaUserService) Jscode2Session(jscode string) (*response.JsCode2SessionResult, error) {
	config := user.wxaService.GetWxaConfig()
	var result response.JsCode2SessionResult
	params := map[string]interface{}{
		"appid":      config.AppId,
		"secret":     config.Secret,
		"js_code":    jscode,
		"grant_type": "authorization_code",
	}
	err := user.wxaService.Get(jsocdeToSessionUrl, &params, &result)

	if err != nil {
		err = errors.New("Failed to get jscode session" + err.Error())
		log.Printf("Failed to get jscode session：%s", err.Error())
		return nil, err
	}

	if result.ErrCode != 0 {
		err = errors.New(result.ErrMsg)
		log.Printf(result.ErrMsg)
		return nil, err
	}

	return &result, nil
}

// 用户支付完成后，获取该用户的 UnionId，无需用户授权
func (user *wxaUserService) GetPaidUnionIdByTransactionId(openId, transactionId string) (unionid string, err error) {

	if common.IsBlank(openId) {
		return "", errors.New("openId is blank")
	} else if common.IsBlank(transactionId) {
		return "", errors.New("transactionId is blank")
	}
	var params = map[string]interface{}{
		"openid":         openId,
		"transaction_id": transactionId,
	}
	var result response.GetPaidUnionId
	err = user.wxaService.Get(getPaidUnionId, &params, &result)

	if err != nil {
		err = errors.New("failed to get unionid:" + err.Error())
		log.Printf("failed to get unionid：%s", err.Error())
		return "", err
	}

	if result.ErrCode != 0 {
		err = errors.New(result.ErrMsg)
		log.Printf(result.ErrMsg)
		return "", err
	}
	return result.Unionid, nil
}

// 用户支付完成后，获取该用户的 UnionId，无需用户授权
func (user *wxaUserService) GetPaidUnionIdByMchIdAndOutTradeNo(openId, mchId, outTradeNo string) (unionid string, err error) {

	if common.IsBlank(openId) {
		return "", errors.New("openId is blank")
	} else if common.IsBlank(mchId) {
		return "", errors.New("mchId is blank")
	} else if common.IsBlank(outTradeNo) {
		return "", errors.New("outTradeNo is blank")
	}
	var params = map[string]interface{}{
		"openid":       openId,
		"mch_id":       mchId,
		"out_trade_no": outTradeNo,
	}
	var result response.GetPaidUnionId
	err = user.wxaService.Get(getPaidUnionId, &params, &result)

	if err != nil {
		err = errors.New("failed to get unionid" + err.Error())
		log.Printf("failed to get unionid：%s", err.Error())
		return "", err
	}

	if result.ErrCode != 0 {
		err = errors.New(result.ErrMsg)
		log.Printf(result.ErrMsg)
		return "", err
	}
	return result.Unionid, nil
}
