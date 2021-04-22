package models

type WxError struct {
	// 微信错误码
	ErrCode int32 `json:"errcode"`
	// 微信错误信息
	ErrMsg string `json:"errmsg"`
}
