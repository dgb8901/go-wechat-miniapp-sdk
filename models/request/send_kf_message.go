package request

// SendKfMessage 发送客服消息请求参数
type SendKfMessage struct {
	// 用户的 OpenID
	ToUser string `json:"touser"`
	// 消息类型
	MsgType string `json:"msgtype"`
	// 文本消息，msgtype="text" 时必填
	Text Text `json:"text"`
	// 图片消息，msgtype="image" 时必填
	Image Image `json:"image"`
	// 图文链接，msgtype="link" 时必填
	Link Link `json:"link"`
	// 小程序卡片，msgtype="miniprogrampage" 时必填
	MiniProgramPage MiniProgramPage `json:"miniprogrampage"`
}

// Text 的结构
type Text struct {
	// 文本消息内容
	Content string `json:"content"`
}

// Image 的结构
type Image struct {
	// 发送的图片的媒体ID，通过 新增素材接口 上传图片文件获得。
	MediaId string `json:"media_id"`
}

// Link 的结构
type Link struct {
	// 消息标题
	Title string `json:"title"`
	// 图文链接消息
	Description string `json:"description"`
	// 图文链接消息被点击后跳转的链接
	Url string `json:"url"`
	// 图文链接消息的图片链接，支持 JPG、PNG 格式，
	// 较好的效果为大图 640 X 320，小图 80 X 80
	ThumbUrl string `json:"thumb_url"`
}

// MiniProgramPage 的结构
type MiniProgramPage struct {
	// 消息标题
	Title string `json:"title"`
	// 小程序的页面路径，跟app.json对齐，支持参数，
	// 比如pages/index/index?foo=bar
	PagePath string `json:"pagepath"`
	// 小程序消息卡片的封面， image 类型的 media_id，
	// 通过 新增素材接口 上传图片文件获得，建议大小为 520*416
	ThumbMediaId string `json:"thumb_media_id"`
}
