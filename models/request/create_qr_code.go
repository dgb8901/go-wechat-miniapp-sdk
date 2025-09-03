package request

// CreateQrCode 获取小程序二维码请求参数
type CreateQrCode struct {
	// 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；
	// 对于小游戏，可以只传入 query 部分，来实现传参效果，
	// 如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	Path string `json:"path"`
	// 二维码的宽度，单位 px。最小 280px，最大 1280px
	Width int `json:"width"`
}
