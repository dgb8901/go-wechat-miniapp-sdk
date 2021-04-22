package request

// 下发客服当前输入状态请求参数
type SetTyping struct {
	// 用户的 OpenID
	ToUser string `json:"touser"`
	// command 的合法值: Typing,CancelTyping
	Command string `json:"command"`
}
