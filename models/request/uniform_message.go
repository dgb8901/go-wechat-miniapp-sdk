package request

// 统一服务消息请求参数
type UniformMessage struct {
	// 用户openid，可以是小程序的openid，
	// 也可以是mp_template_msg.appid对应的公众号的openid
	ToUser string `json:"touser"`
	// 小程序模板消息相关的信息，可以参考小程序模板消息接口;
	// 有此节点则优先发送小程序模板消息
	WeappTemplateMsg *WeappTemplateMsg `json:"weapp_template_msg"`
	//公众号模板消息相关的信息，可以参考公众号模板消息接口；
	//有此节点并且没有weapp_template_msg节点时，发送公众号模板消息
	MpTemplateMsg *MpTemplateMsg `json:"mp_template_msg"`
}

// weapp_template_msg 的结构
type WeappTemplateMsg struct {
	// 小程序模板ID
	TemplateId string `json:"template_id"`
	// 小程序页面路径
	Page string `json:"page"`
	// 小程序模板消息formid
	FormId string `json:"form_id"`
	// 小程序模板数据
	Data string `json:"data"`
	// 小程序模板放大关键词
	EmphasisKeyword string `json:"emphasis_keyword"`
}

// mp_template_msg 的结构
type MpTemplateMsg struct {
	// 公众号appid，要求与小程序有绑定且同主体
	Appid string `json:"appid"`
	// 公众号模板id
	TemplateId string `json:"template_id"`
	// 公众号模板消息所要跳转的url
	Url string `json:"url"`
	// 公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
	Miniprogram string `json:"miniprogram"`
	// 公众号模板消息的数据
	Data string `json:"data"`
}
