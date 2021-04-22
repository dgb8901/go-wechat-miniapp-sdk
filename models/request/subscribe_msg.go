package request

// 订阅消息请求参数
type SubscribeMsg struct {
	// 接收者（用户）的 openid
	ToUser string `json:"touser"`
	// 所需下发的订阅模板id
	TemplateId string `json:"template_id"`
	// 点击模板卡片后的跳转页面，仅限本小程序内的页面。
	// 支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Page string `json:"page"`
	// 模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
	Data interface{} `json:"data"`
	// 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	MiniprogramState string `json:"miniprogram_state"`
	// 进入小程序查看”的语言类型，
	// 支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
	Lang string `json:"lang"`
}
